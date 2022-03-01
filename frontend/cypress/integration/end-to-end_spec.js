
//=====================End to End ======================//
describe('Select com tag', () =>{
    before(() => {
        cy.visit("http://localhost:8080").wait(1000)

    })
    it('should com and click next btn', () =>{
        cy.get('button').contains('com').click().wait(1000)   
        cy.get('button').contains('food').click().wait(1000)
        cy.get('button').contains('it').click().wait(1000)
        cy.get('button').contains('com').click().wait(1000)
        cy.get('button').contains('food').click().wait(1000)
        cy.get('#next-button').click().wait(1000)
        cy.get('#productName').contains("Moblie phone 2 hand").wait(1000)
        cy.contains('com').click().wait(1000)
        cy.contains('it').click().wait(1000)
        cy.contains('food').click().wait(1000)
        cy.get('input').type('thai').wait(1000)
        cy.get('#search').click().wait(1000)
        
    })
})
// ======================================================//