proxy
=====

Use this if you're sick of typing things like

    $ export NO_PROXY=$NO_PROXY,something.local
    $ export https_proxy=$http_proxy

Install with:

    $ go install github.com/kragniz/proxy

Use it:

    $ proxy on
    $ proxy off
    $ proxy ignore localhost
    $ proxy ignore something.local

Set the default proxy:

    $ echo 'ProxyServer: "proxy.dootcorp.com"' > ~/.proxy.yaml
