package commands

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dictyBase/gh-issue/auth"
	"github.com/dictyBase/gh-issue/resources"
	"gopkg.in/urfave/cli.v1"
)

func RunServer(c *cli.Context) {
	/*if err := ValidateServerOptions(c); err != nil {
		log.Fatal(err)
	}
	var logMw *middlewares.Logger
	if c.IsSet("log") {
		w, err := os.Create(c.String("log"))
		if err != nil {
			log.Fatalf("cannot open log file %q\n", err)
		}
		defer w.Close()
		logMw = middlewares.NewFileLogger(w)
	} else {
		logMw = middlewares.NewLogger()
	}
	gmClient, err := auth.GetGmailClient(c)
	if err != nil {
		log.Fatal(err)
	}*/
	ghClient, err := auth.GetGithubClient(c)
	if err != nil {
		log.Fatal(err)
	}
	//r := routes.NewRouter()
	mux := http.NewServeMux()
	/*valMw := &middlewares.GmailSubscription{
		fmt.Sprintf(
			"projects/%s/subscriptions/%s",
			c.String("project"),
			c.String("subscription"),
		),
	}
	hdb, err := history.NewHistoryDb(
		fmt.Sprintf(
			"%s:%d",
			c.String("redis-address"),
			c.Int("redis-port"),
		),
	)
	if err != nil {
		log.Fatalf("error in connecting to history db %s\n", err)
	}
	/*
		lm := labels.NewLabelManager(gmClient)
		err = lm.GenerateCache()
		if err != nil {
			log.Fatalf("error in generating labels cache %s\n", err)
		}
		if !lm.HasLabel(c.String("label")) {
			log.Fatalf("given label %s does not exist\n", c.String("label"))
		}

		logger := log.New(os.Stderr, "gmail-webhook", log.Lshortfile)
		if c.IsSet("app-log") {
			l, err := os.Create(c.String("app-log"))
			if err != nil {
				log.Fatalf("error creating log file %s\n", err)
			}
			defer l.Close()
			logger = log.New(l, "gmail-webhook", log.Lshortfile)
		}
		rgxp, err := regexp.Compile(`Order_Type:(\w+)\|(\w+)`)
		if err != nil {
			log.Fatalf("error in creating regexp %s\n", err)
		}*/
	dsc := &handlers.Client{

		Github: ghClient,
		//Label:      lm.Name2Id(c.String("label")),
		Repository: c.String("repository"),
		Owner:      c.String("owner"),

		//Logger: logger,
	}
	/*dscChain := apollo.New(
		apollo.Wrap(logMw.LoggerMiddleware),
		middlewares.DecodeMiddleware,
		valMw.ValidateMiddleware,
	).With(context.Background()).ThenFunc(dsc.StockOrderHandler)
	*/
	//r.Post("/order", handlers.GithubHandler())

	mux.HandleFunc("/gh/order", dsc.GithubHandler)
	log.Printf("Starting web server on port %d\n", c.Int("port"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Int("port")), mux))
}
