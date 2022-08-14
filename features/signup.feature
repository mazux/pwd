Feature: Signup with a new profile

    All possible scenarios are:
        - Error path: Profile already exist
        - Happy path: Profile got created with no problems
    
    Scenario: Profile already exist
        Given I have a profile for username "MAZux"
        When I signup using username "MAZux"
        Then I will get this error "profile already exists"

    Scenario: Profile got created with no problems
        Given I have no profiles stored
        When I signup using username "MAZux"
        Then I will have a stored profile with username "MAZux" and empty list of logins
