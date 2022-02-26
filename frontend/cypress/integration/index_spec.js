describe('Select all tag', () =>{
    before(() => {
        cy.visit("http://localhost:8080").wait(1000)

    })
    it('should display all tag display', () =>{
        cy.get('button').click({ multiple: true }).wait(1000)
        cy.get('#next-button').click().wait(1000)
        
    })
})
describe('Select com tag', () =>{
    before(() => {
        cy.visit("http://localhost:8080").wait(1000)

    })
    it('should com and click next btn', () =>{
        cy.get('button').contains('com').click().wait(1000)
        cy.get('#next-button').click().wait(1000)
        cy.get('#productName').contains("Computer 2").wait(1000)
    })
})
describe('Select com tag and undo click tag', ()=>{
    before(() => {
        cy.visit("http://localhost:8080").wait(1000)

    })
    it('should undo tag', () =>{
        cy.get('button').contains('com').click().wait(1000)
        cy.get('button').contains('com').click().wait(1000)
    })
})
describe('Select com tag and undo click tag and select food tag', ()=>{
    before(() => {
        cy.visit("http://localhost:8080").wait(1000)

    })
    it('should undo com tag and select food tag', () =>{
        cy.get('button').contains('com').click().wait(1000)
        cy.get('button').contains('com').click().wait(1000)
        cy.get('button').contains('food').click().wait(1000)
        cy.get('#next-button').click().wait(1000)
    })
})