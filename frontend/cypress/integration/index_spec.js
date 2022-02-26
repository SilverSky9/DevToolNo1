describe('Select tag', () =>{
    before(() => {
        cy.visit("http://localhost:8080").wait(1000)

    })
    it('should display all tag display', () =>{
        // cy.get('button').click({ multiple: true }).wait(500)
        cy.get('button').contains('com').click().wait(1000)
        cy.get('#next-button').click()
    })
})