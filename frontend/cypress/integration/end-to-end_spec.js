

describe('End-to-end', () =>{
    before(() => {
        cy.visit("http://localhost:8080").wait(1000)
    })
    it('should select and deselect tags of product', () =>{
        cy.get('button').contains('ความสวยความงาม').click().wait(1000)   
        cy.get('button').contains('อาหาร').click().wait(1000)
        cy.get('button').contains('สัตว์เลี้ยง').click().wait(1000)
        cy.get('button').contains('ไอที').click().wait(1000)

        cy.get('button').contains('ความสวยความงาม').click().wait(1000)   
        cy.get('button').contains('อาหาร').click().wait(1000)
        cy.get('button').contains('สัตว์เลี้ยง').click().wait(1000)
        cy.get('button').contains('ไอที').click().wait(1000)

        cy.get('button').contains('ความสวยความงาม').click().wait(1000)  
        cy.get('button').contains('สัตว์เลี้ยง').click().wait(1000) 
  
    })

    it('should select next btn', () =>{
        cy.get('#next-button').click().wait(1000)
    })

    it('should search name of product', () =>{
        cy.get('input').type('นก').wait(1000)
        cy.get('#search').click().wait(1000)
    })

    it('should filter product by tag', () =>{
        cy.get('div').contains('ความสวยความงาม').click().wait(1000)   
        cy.get('div').contains('อาหาร').click().wait(1000)
        cy.get('div').contains('สัตว์เลี้ยง').click().wait(1000)
        cy.get('div').contains('ไอที').click().wait(1000)

    })

    it('should create post', () =>{

        cy.get('button').contains('+').click().wait(1000)
        cy.get('button').contains('Cancle').click().wait(1000) 
        cy.get('button').contains('+').click().wait(1000)
        cy.get('#productNameInput').type('มาม่ารสต้มยำ').click().wait(1000)
        cy.get('#productDetail').type('เผ็ดมาก ๆ รสนี้').click().wait(1000)
        cy.get('#phone').type('0824919243').click().wait(1000)

        cy.get('button').contains('Post').click().wait(1000) 
        
        cy.get('#price').type('20').click().wait(1000)
        cy.get('#amount').type('100').click().wait(1000)
        cy.get('select#location').select('1').should('have.value', '1').wait(1000)
        cy.get('select#location').select('2').should('have.value', '2').wait(1000)
        cy.get('select#location').select('3').should('have.value', '3').wait(1000)
        cy.get('select#tagName').select('อาหาร').should('have.value', 'อาหาร').wait(1000)
        cy.get('button').contains('Post').click().wait(1000) 
        
    })


  





})
