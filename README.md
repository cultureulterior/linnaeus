linnaeus
========

Linnaeus serializes machine-local data to json files. These tools are meant to be used with consul-template, to work as a replacement for 'facter' in the consul infrastructure. 
With consul-template and linnaeus, you can do things like:
````
{{with $d := file "/tmp/aws.json" | parseJSON}}
stats show-desc {{$d.instance_id}}
{{end}}
````
The project is named for the Swedish scientist Linnaeus, who classified the natural world.
