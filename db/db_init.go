package db

import (
    "log"
    "github.com/go-pg/pg/v10"
    "github.com/go-pg/migrations/v8"

)

func StartDB() (*pg.DB, error) {
   
    opts := &pg.Options{
      Addr: "db:5432",
      User: "postgres",
      Password: "admin",
    }

    db := pg.Connect(opts)

    collection := migrations.NewCollection()
    err := collection.DiscoverSQLMigrations("sql/init")
    if err != nil {
        return nil, err
    }

    _, _, err = collection.Run(db, "init")
    if err != nil {
        return nil, err
    }

    oldVersion, newVersion, err := collection.Run(db, "up")
    if err != nil {
        return nil, err
    }

    if newVersion != oldVersion {
      log.Printf("Migrated from %d version to %d version", oldVersion, newVersion)
    } else {
      log.Printf("Version is %d", newVersion)
    }

    return db, err
}
