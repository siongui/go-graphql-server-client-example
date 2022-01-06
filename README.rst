================================
Go GraphQL Server/Client Example
================================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/go-graphql-server-client-example?status.svg
   :target: https://godoc.org/github.com/siongui/go-graphql-server-client-example

.. image:: https://github.com/siongui/go-graphql-server-client-example/workflows/ci/badge.svg
    :target: https://github.com/siongui/go-graphql-server-client-example/blob/master/.github/workflows/ci.yml

.. image:: https://goreportcard.com/badge/github.com/siongui/go-graphql-server-client-example
   :target: https://goreportcard.com/report/github.com/siongui/go-graphql-server-client-example

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://github.com/siongui/go-graphql-server-client-example/blob/master/UNLICENSE


GraphQL_ server/client example for Go_.

Development Environment:

  - `Ubuntu 20.04`_
  - `Go 1.17`_


Usage
+++++

After git clone this repo, generate and run server:

.. code-block:: bash

  $ cd /path/to/this/repo/
  $ cd server/
  # Install github.com/99designs/gqlgen
  $ make install
  # go generate server code
  $ make generate
  # Run server
  $ make

After server is running, you use curl to send test request to server:

.. code-block:: bash

  $ make curl

You can also use client to send test request to server:

.. code-block:: bash

  $ cd /path/to/this/repo/
  $ cd client/
  $ make

See Makefile in ``server/`` and ``client/`` directories for actual commands.


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] `github.com/99designs/gqlgen <https://github.com/99designs/gqlgen>`_
.. [2] `github.com/machinebox/graphql <https://github.com/machinebox/graphql>`_
.. [3] `go - What is the correct shape of a curl POST request to a gqlgen GraphQL API? - Stack Overflow <https://stackoverflow.com/questions/54271405/what-is-the-correct-shape-of-a-curl-post-request-to-a-gqlgen-graphql-api>`_

.. _Go: https://golang.org/
.. _GraphQL: https://graphql.org/
.. _Ubuntu 20.04: https://releases.ubuntu.com/20.04/
.. _Go 1.17: https://golang.org/dl/
.. _UNLICENSE: https://unlicense.org/
