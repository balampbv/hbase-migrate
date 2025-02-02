# HBase Migration Library

This Go library is designed to help manage HBase schema migrations, supporting both upward and downward migrations. It integrates with HBase via Thrift to apply changes like creating tables, deleting tables, and modifying column families.

## Table of Contents

- [Key Features](#key-features)
- [HBase Integration via Thrift](#hbase-integration-via-thrift)
    - [Steps to Generate Thrift Code](#steps-to-generate-thrift-code)

### Key Features

- **Table Creation**: Create new tables with specified column families.
- **Table Deletion**: Delete existing tables.
- **Modify Column Families**: Add or modify column families for tables.
- **Migration Management**: Versioned migration handling, ensuring migrations are applied in order.

---

## HBase Integration via Thrift

This project integrates with HBase using Thrift for schema management. The `hbase.thrift` file defines the RPC services and data structures used to interact with HBase.

### Steps to Generate Thrift Code

To generate Go code from the `hbase.thrift` file, follow these steps:

1. **Install Thrift Compiler**

   If you don’t have the Thrift compiler installed, you can install it as follows:

    - **Ubuntu**:
      ```bash
      sudo apt-get install thrift-compiler
      ```

    - **macOS (Homebrew)**:
      ```bash
      brew install thrift
      ```

2. **Generate Go Code**

   Once the Thrift compiler is installed, run the following command to generate Go code from the `hbase.thrift` file:

   ```bash
    cd hbase
    thrift --gen go hbase.thrift
