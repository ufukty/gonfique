**Todo**

-   reduce flag combinations to ease up development, simplify process pipeline

-   always use organize behavior, assign smarter names by default without exporting to outside package

-   redesign array support

    -   on compatible item types; define common props in a distinct struct type then embed it.

    -   on item type conflict; define array type as interface, and comply the interface on each item type.

-   type resolution for `time.Duration`, `uuid` etc.

-   redesign instruction file syntax:

    ```yml
    type-names:
        name: Name
        services.**.name: ServiceName

    array-type-:
        "**.services": combine | interface
    ```

    ```yml
    name:
        typename: Name

    **.grace-period:
        type: time.Duration

    services.**.name: 
        typename: ServiceName
        type: string
        assert: compatible
    ```