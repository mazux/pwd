Feature: Add login credentials to an existed profile

    To cover:
    - Adding a new login to a profile:
        - Login is already there
        - Login is successfully added
    - Fetch a login from a profile
        - Login is not found
    - Remove a login from a profile

    Background: Clear the storage
        Given I have empty storage

    Scenario: Adding a duplicated login
        Given I have a profile for username "MAZux"
        And It has a login for domain "stackoverflow.com" and username "MAZux7"
        When I add a new login for domain "stackoverflow.com" and username "MAZux7"
        Then I will get this error "login already exists in profile"
