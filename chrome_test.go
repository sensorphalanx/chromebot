package chromebot_test

import (
	"context"
	"log"
	"testing"
	"time"
	"flag"
	"os/exec"
	"strings"
	"strconv"

	"github.com/chromedp/chromedp"
	"github.com/nanitefactory/chromebot"
)

var bFlagHeadless = false

func init() {
	flag.BoolVar(&bFlagHeadless, "headless", false, "If set to true, the browser runs in headless mode.")
	{
		major, minor, _, err := func() (major int, minor int, patch int, err error) {
			cmd := exec.Command("go", "version")
			out, err := cmd.Output()
			if err != nil {
				return
			}
			semantic := strings.Split(strings.Fields(strings.TrimLeft(string(out), "go version go"))[0], ".")
			patch, err = strconv.Atoi(semantic[2])
			minor, err = strconv.Atoi(semantic[1])
			major, err = strconv.Atoi(semantic[0])
			return
		}() // getGoVersion
		if err == nil && (major < 1 || (major == 1 && minor < 13)) {
			Init() // if 1.12 or below for sure
		}
	}
}

func Init() {
	flag.Parse()
}

func TestChromedp(t *testing.T) {
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), chromedp.Headless)
	defer cancel()
	ctx2, _ := chromedp.NewContext(ctx)
	chromedp.Run(ctx2)
}

func TestNewChrome(t *testing.T) {
	if !bFlagHeadless {
		chromebot.New(bFlagHeadless).Close()
	}
	chromebot.New(true).Close()
}

func TestChromeManyTabs(t *testing.T) {
	testCountTab := func(c *chromebot.Chrome, n int) {
		if c.CountTabs() != n {
			panic("c.CountTabs() not working well")
		}
	}
	c := chromebot.New(bFlagHeadless)
	defer c.Close()
	//
	n := 1
	testCountTab(c, n)
	//
	tab1 := c.AddNewTab()
	n++
	testCountTab(c, n)
	//
	tab2 := c.AddNewTab()
	n++
	testCountTab(c, n)
	//
	for i := 0; i < 10; i++ {
		c.AddNewTab()
		n++
		testCountTab(c, n)
	}
	//
	tab2.Close()
	tab1.Close()
	for i := 0; i < c.CountTabs(); i++ {
		c.Tab(i).Close()
		time.Sleep(time.Millisecond * 100)
	}
	testCountTab(c, n)
}

func TestChromeClose(t *testing.T) {
	c1 := chromebot.New(bFlagHeadless)
	defer c1.Close()
	defer c1.Close()
	defer c1.Close()
	c1.AddNewTab()
	c1.Tab(0).Close()
	c1.Tab(0).Close()
	c1.Tab(0).Close()
	c1.AddNewTab()
}

func TestChromeTabAsBrowser(t *testing.T) {
	defer func() { // expecting a panic
		if r := recover(); r == nil {
			panic("must panic but it didn't panic")
		}
	}()
	c := chromebot.New(bFlagHeadless)
	defer c.Close()
	c.AddNewTab(chromedp.WithBrowserOption(chromedp.WithBrowserLogf(func(str string, args ...interface{}) {
		log.Println()
	})))
}

func TestDeadChrome(t *testing.T) {
	c := chromebot.New(bFlagHeadless)
	deadTab := c.Tab(0)
	c.Close()
	c.Close()
	c.Close()
	c.Close()
	if tab := c.Tab(0); tab == nil {
		panic(tab)
	}
	if nTabs := c.CountTabs(); nTabs != 1 {
		panic(nTabs)
	}
	c.Tab(0).Close()
	if c.Tab(0) != deadTab {
		panic("failed to get the first tab")
	}
	if tab := c.AddNewTab(); tab != nil {
		panic(tab)
	}
	if err := c.Listen(func(ev interface{}) {}); err == nil {
		panic("Listen() should error")
	}
	select {
	case <-c.Dead():
	default:
		panic("Dead() not working well")
	}
}
