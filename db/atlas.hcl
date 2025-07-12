atlas {
  lint {
    destructive = true
  }
  migrate {
    dir = "file://db/migrations"
  }
}

env "local" {
  url = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
  dev = "docker://postgres/16"
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

