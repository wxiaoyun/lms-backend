#!/bin/sh
chown -R app:app /app/file_storage
chmod -R 766 /app/file_storage
exec "$@"
