// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.906
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Login() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<html><head><title>Login</title><script src=\"/static/js/htmx.min.js\"></script><style>\r\n\t\t\t\tbody, html {\r\n\t\t\t\t\tmargin: 0;\r\n\t\t\t\t\tpadding: 0;\r\n\t\t\t\t\theight: 100%;\r\n\t\t\t\t\tfont-family: sans-serif;\r\n\t\t\t\t\tbackground-color: #f9f9f9;\r\n\t\t\t\t}\r\n\r\n\t\t\t\t.container {\r\n\t\t\t\t\tdisplay: flex;\r\n\t\t\t\t\talign-items: center;\r\n\t\t\t\t\tjustify-content: center;\r\n\t\t\t\t\theight: 100%;\r\n\t\t\t\t}\r\n\r\n\t\t\t\t.login-box {\r\n\t\t\t\t\tbackground: white;\r\n\t\t\t\t\tpadding: 2rem;\r\n\t\t\t\t\tborder-radius: 8px;\r\n\t\t\t\t\tbox-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);\r\n\t\t\t\t\twidth: 300px;\r\n\t\t\t\t}\r\n\r\n\t\t\t\t.login-box h1 {\r\n\t\t\t\t\ttext-align: center;\r\n\t\t\t\t\tmargin-bottom: 1.5rem;\r\n\t\t\t\t}\r\n\r\n\t\t\t\tform {\r\n\t\t\t\t\tdisplay: grid;\r\n\t\t\t\t\tgap: 1rem;\r\n\t\t\t\t}\r\n\r\n\t\t\t\tinput {\r\n\t\t\t\t\tpadding: 0.5rem;\r\n\t\t\t\t\tborder: 1px solid #ccc;\r\n\t\t\t\t\tborder-radius: 4px;\r\n\t\t\t\t\tfont-size: 1rem;\r\n\t\t\t\t}\r\n\r\n\t\t\t\tbutton {\r\n\t\t\t\t\tbackground-color: #313131;\r\n\t\t\t\t\tcolor: white;\r\n\t\t\t\t\tborder: none;\r\n\t\t\t\t\tpadding: 0.5rem;\r\n\t\t\t\t\tborder-radius: 4px;\r\n\t\t\t\t\tfont-size: 1rem;\r\n\t\t\t\t\tcursor: pointer;\r\n\t\t\t\t}\r\n\r\n\t\t\t\tbutton:hover {\r\n\t\t\t\t\tbackground-color: #0056b3;\r\n\t\t\t\t}\r\n\r\n\t\t\t\t#response {\r\n\t\t\t\t\tmargin-top: 0.5rem;\r\n\t\t\t\t\ttext-align: center;\r\n\t\t\t\t}\r\n\r\n\t\t\t\t.signup-link {\r\n\t\t\t\t\ttext-align: center;\r\n\t\t\t\t\tmargin-top: 1rem;\r\n\t\t\t\t\tfont-size: 0.9rem;\r\n\t\t\t\t}\r\n\t\t\t</style></head><body><div class=\"container\"><div class=\"login-box\"><h1>Login</h1><form id=\"login\"><input type=\"email\" name=\"email\" placeholder=\"Email (@example.com)\" pattern=\"^[a-zA-Z0-9._%+-]+@example\\.com$\" required> <input type=\"password\" name=\"password\" placeholder=\"Password\" pattern=\"^(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d]{8,}$\" minlength=\"8\" required> <button type=\"submit\">Login</button></form><div id=\"response\"></div><div class=\"signup-link\"><p>Need to create an account? <a href=\"/signup\">Sign up</a></p></div></div></div><script>\r\n\t\t\t\tdocument.addEventListener(\"DOMContentLoaded\", function () {\r\n\t\t\t\t\tconst form = document.getElementById(\"login\");\r\n\t\t\t\t\tconst response = document.getElementById(\"response\");\r\n\r\n\t\t\t\t\tform.addEventListener(\"submit\", function (event) {\r\n\t\t\t\t\t\tevent.preventDefault();\r\n\r\n\t\t\t\t\t\tif (!form.checkValidity()) {\r\n\t\t\t\t\t\t\tresponse.innerHTML = \"<p style='color:red;'>Please fill out the form correctly.</p>\";\r\n\t\t\t\t\t\t\treturn;\r\n\t\t\t\t\t\t}\r\n\r\n\t\t\t\t\t\tconst formData = new FormData(form);\r\n\r\n\t\t\t\t\t\thtmx.ajax(\"POST\", \"/login\", {\r\n\t\t\t\t\t\t\ttarget: \"#response\",\r\n\t\t\t\t\t\t\tswap: \"innerHTML\",\r\n\t\t\t\t\t\t\tvalues: Object.fromEntries(formData.entries())\r\n\t\t\t\t\t\t});\r\n\t\t\t\t\t});\r\n\t\t\t\t});\r\n\t\t\t</script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
