/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"keycloak-demo-5-client/config"
	"keycloak-demo-5-client/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start server",
	Long:  `start server`,
	Run: func(cmd *cobra.Command, args []string) {

		config := config.LoadConfig()
		fmt.Printf("%+v", config)
		h := handler.Handler{Config: config}

		r := chi.NewRouter()
		r.Use(middleware.Logger)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`
<body>
	<section>
		<div><a href="/auth/start?scope=openid profile readdata">scope: openid, profile, readdata<a/></div>
		<div><a href="/auth/start?scope=profile readdata">scope: profile, readdata<a/></div>
		<div><a href="/auth/start?scope=readdata">scope: readdata<a/></div>
		<div><a href="/auth/start?scope=profile">scope: profile<a/></div>
	</section>
</body>
`))
		})
		r.Route("/auth", func(r chi.Router) {
			r.Get("/start", h.StartAuth)
			r.Get("/callback", h.Callback)
			r.Get("/introspect", h.Introspect)
		})
		r.Route("/user", func(r chi.Router) {
			r.Get("/profile", h.Profile)
		})

		http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), r)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
