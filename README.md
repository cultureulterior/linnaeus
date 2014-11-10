linnaeus
========

Linnaeus serializes machine-local data to json files. These tools are meant to be used with https://github.com/hashicorp/consul-template, to work as a replacement for 'facter' in the consul infrastructure. 
With consul-template and linnaeus, you can do things like:
(in `haproxy.ctmpl`)
````
{{with $d := file "/tmp/aws.json" | parseJSON}}
stats show-desc {{$d.instance_id}}
{{end}}
````
### Aws2json

This tool serializes the tags and some other data (region, az, type) of an aws node.

### Unix2json

This tool serializes the amount of RAM and number of CPUs of a host. More features welcome!

The project is named for the Swedish scientist Linnaeus, who classified the natural world.
