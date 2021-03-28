#!/bin/bash
docker build -t abelce/product_db:1.0.0 .
docker push abelce/product_db:1.0.0
