# Reloader

usage: reloader "path" "command" &
Watches "path". Executes "command" on changes.

##example:
reload page in firefox (with <a href='https://addons.mozilla.org/en-US/firefox/addon/remote-control/'>remote control plugin</a>)
<pre>reloader /path/to/project "echo reload | nc localhost 3200" &</pre>

