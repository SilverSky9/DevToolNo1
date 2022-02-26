describe('Search by title', () =>{
    before(() =>{
        cy.visit("http://localhost:8080/Home")
    })
    it('search "com"', () =>{
        cy.get('input').type('com')
        cy.get('#search').click()
    })
})