#!/bin/bash
docker build -t abelce/user_db:1.0.0 .
docker push abelce/user_db:1.0.0
