# Reloader
Filesystem watcher written in go (golang)

##Usage: 
<pre>reloader "path" "command" &</pre>
<p>
Watches "path". Executes "command" on changes.
</p>

##Example:
Reload page in firefox (with <a href='https://addons.mozilla.org/en-US/firefox/addon/remote-control/'>remote control plugin</a>)
<pre>reloader /path/to/project "echo reload | nc localhost 32000" &</pre>

