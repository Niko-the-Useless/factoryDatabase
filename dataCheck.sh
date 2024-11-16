#!/bin/zsh

DB_NAME="factorydb"

echo "tables: "
sqlite3 $DB_NAME <<EOF
.tables
EOF

echo "table to check: "
read TABLE_NAME

sqlite3 $DB_NAME <<EOF
.headers on
.mode column
.nullvalue null
SELECT * FROM $TABLE_NAME;
EOF
