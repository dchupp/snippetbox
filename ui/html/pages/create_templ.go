// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package ui

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/dchupp/snippetbox/ui"

func createSnippet(data ui.TemplateData) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

// {{define "title"}}Create a New Snippet{{end}}

// {{define "main"}}
// <form action='/snippet/create' method='POST'>
//     <!-- Include the CSRF token -->
//     <input type='hidden' name='csrf_token' value='{{.CSRFToken}}'>
//     <div>
//         <label>Title:</label>
//         {{with .Form.FieldErrors.title}}
//             <label class='error'>{{.}}</label>
//         {{end}}
//         <input type='text' name='title' value='{{.Form.Title}}'/>
//     </div>
//     <div>
//         <label>Content:</label>
//         {{with .Form.FieldErrors.content}}
//             <label class='error'>{{.}}</label>
//         {{end}}
//         <textarea name='content'>{{.Form.Content}}</textarea>
//     </div>
//     <div>
//         <label>Delete in:</label>
//         {{with .Form.FieldErrors.expires}}
//             <label class='error'>{{.}}</label>
//         {{end}}
//         <input type='radio' name='expires' value='365' {{if (eq .Form.Expires 365)}}checked{{end}}> One Year
//         <input type='radio' name='expires' value='7' {{if (eq .Form.Expires 7)}}checked{{end}}> One Week
//         <input type='radio' name='expires' value='1' {{if (eq .Form.Expires 1)}}checked{{end}}> One Day
//     </div>
//     <div>
//         <input type='submit' value='Publish snippet'>
//     </div>
// </form>
// {{end}}
