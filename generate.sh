#!/bin/bash

# //go:generate statik -src=precure/rubicure/config/girls -include=*.yml -p=precure
statik -src=precure/rubicure/config/girls -include=*.yml -p=precure
