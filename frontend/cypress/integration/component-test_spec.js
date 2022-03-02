describe('Search by title', () =>{
    before(() =>{
        cy.visit("http://localhost:8080/Home")
    })
    it('search "com"', () =>{
        cy.get('input').type('com').wait(1000)
        cy.get('#search').click().wait(1000)
    })

})
describe('Select tag in home page', () =>{
    before(() =>{
        cy.visit("http://localhost:8080/Home")
    })
    it('Select all tag in home page', () =>{
        cy.get('.tag').eq(0).click().wait(1000)
        cy.get('.tag').eq(1).click().wait(1000)
        cy.get('.tag').eq(2).click().wait(1000)
    })
})

describe('Search by blank space', () =>{
    before(() =>{
        cy.visit("http://localhost:8080/Home")
    })
    it('search ""', () =>{
        cy.get('#search').click().wait(1000)
    })
})

describe('Search by white space', () =>{
    before(() =>{
        cy.visit("http://localhost:8080/Home")
    })
    it('search " "', () =>{
        cy.get('#search').click().wait(1000)
    })
})