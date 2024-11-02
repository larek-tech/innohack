// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Chat() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col w-full max-w-full sm:max-w-lg md:max-w-2xl lg:max-w-3xl xl:max-w-4xl bg-white border border-gray-300 rounded-lg shadow-md h-3/4 min-h-[500px]\" hx-ext=\"ws\" ws-connect=\"/chat/ws\"><!-- Chat messages container with flex-grow to take available space --><div id=\"notifications\" class=\"flex-1 p-3 sm:p-4 overflow-y-auto space-y-4\" hx-swap-oob=\"beforeend\"><!-- Plain text messages will appear here --></div><!-- Chat input form sticks to the bottom of the page --><form id=\"chatForm\" class=\"flex items-center border-t border-gray-300 p-3 sm:p-4\" ws-send=\"submit\" hx-trigger=\"submit\"><input type=\"text\" name=\"content\" class=\"flex-1 p-2 border border-gray-300 rounded-full focus:outline-none focus:border-blue-500\" placeholder=\"Type your message here\" required> <button type=\"submit\" class=\"ml-2 sm:ml-3 px-3 sm:px-4 py-1 sm:py-2 bg-blue-500 text-white rounded-full hover:bg-blue-600 focus:outline-none\">Send</button></form></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate