# Kosha Airtable Connector

Airtable is a spreadsheet-database hybrid, with the features of a database but applied to a spreadsheet. The fields in an Airtable table are similar to cells in a spreadsheet, but have types such as 'checkbox', 'phone number', and 'drop-down list', and can reference file attachments like images.

Users can create a database, set up column types, add records, link tables to one another, collaborate, sort records and publish views to external websites.


![airtable](images/test.png)

Features:

* Supports CRUD operations on Airtable records.
* Supports CRUD operations on Airtable tables.
* Supports CRUD operations on Airtable schema.

The user can use the connector to integrate your data in Airtable with any external system. The API closely follows REST semantics, uses JSON to encode objects, and relies on standard HTTP codes to signal operation outcomes.

The required configuration is your personal access token tied to your Airtable connector and associated scopes.

## OpenAPI Specification

Full connector API specification can be found at [here](openapi.json).


## Authentication

The required configuration is your personal access token tied to your Airtable connector and associated scopes.
