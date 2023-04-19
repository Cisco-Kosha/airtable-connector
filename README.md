# Kosha Airtable Connector

![airtable](images/test.png)

With Airtable you can create spreadsheets that function like databases. 

Using the Kosha Airtable connecter, you can create bases (databases), set up column types, add records, link tables to one another, collaborate, sort records, and integrate your Airtable data with any external system.

The Airtable API follows REST semantics, encodes objects in JSON, and uses standard HTTP response codes.

## Useful Actions 

You can use the Kosha Airtable connector to perform to manage records, tables, fields, and bases.

Refer to the Kosha Airtable connector [API specification](openapi.json) for details.

### Managing Records

A record is an individual unit in a table. Use the Airtable API to create, manage and delete records.

### Managing Fields

Fields are the units in records you use to store custom data. Use the Airtable API to create and update fields.

### Managing Tables

Tables are structural units that enable you to consistently store your Airtable data. 

### Managing Bases

In Airtable a base—or database—is central place to store your data, including tables. You can use the API to list, get, and create bases. 

## Example Usage

The following request creates a table:

```
curl -X POST "https://api.airtable.com/v0/meta/bases/{baseId}/tables" \
-H "Authorization: Bearer YOUR_TOKEN" \
-H "Content-Type: application/json" \
--data '{
    "description": "A to-do list of places to visit",
    "fields": [
      {
        "description": "Name of the apartment",
        "name": "Name",
        "type": "singleLineText"
      },
      {
        "name": "Address",
        "type": "singleLineText"
      },
      {
        "name": "Visited",
        "options": {
          "color": "greenBright",
          "icon": "check"
        },
        "type": "checkbox"
      }
    ],
    "name": "Apartments"
  }'
  ```

## Authentication

Authentication requires your [personal access token](https://airtable.com/developers/web/guides/personal-access-tokens) tied to your Kosha Airtable connector and associated scopes.
