package main

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("Hit the page")
			next.ServeHTTP(w, r)
		})
}
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProd, //as not using https
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

/*
Types of Cookie Modes

1) SameSiteLaxMode
   Example: example.com sets a cookie → {Set-Cookie: session_id=abc123; SameSite=Lax}
   - 🖥️ User types URL in browser ................. ✅ (Allowed)
   - 🔗 User clicks a link from another site (GET) ✅ (Allowed)
   - 📝 Form submission (POST) from another site.. ❌ (Blocked)
   - 🖼️ Image or iframe loaded from another site.. ❌ (Blocked)
   - 🌍 API fetch (fetch() or XMLHttpRequest) ..... ❌ (Blocked)

2) SameSiteStrictMode
   Example: example.com sets a cookie → {Set-Cookie: session_id=abc123; SameSite=Strict}
   - 🖥️ User types URL in browser ................. ✅ (Allowed)
   - 🔗 User clicks a link from another site (GET) ❌ (Blocked)
   - 📝 Form submission (POST) from another site.. ❌ (Blocked)
   - 🖼️ Image or iframe loaded from another site.. ❌ (Blocked)
   - 🌍 API fetch (fetch() or XMLHttpRequest) ..... ❌ (Blocked)

3) SameSiteNoneMode (Requires Secure)
   Example: example.com sets a cookie → {Set-Cookie: session_id=abc123; SameSite=None; Secure}
   - 🖥️ User types URL in browser ................. ✅ (Allowed)
   - 🔗 User clicks a link from another site (GET) ✅ (Allowed)
   - 📝 Form submission (POST) from another site.. ✅ (Allowed)
   - 🖼️ Image or iframe loaded from another site.. ✅ (Allowed)
   - 🌍 API fetch (fetch() or XMLHttpRequest) ..... ✅ (Allowed, if HTTPS)

4) SameSiteDefaultMode (Depends on browser)
   - Behavior varies depending on the browser.
   - Most browsers treat it as `Lax` by default.
*/
