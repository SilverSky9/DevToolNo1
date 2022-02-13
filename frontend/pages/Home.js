import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'
import shopping_cart from '../public/shopping-cart.png' 

const tag_want = []
var data1 = [
    {
        post_id : 1,
        pin : 123,
        product_name : "cat",
        post_date : "12-3-63",
        product_option : "buy",
        price : 520,
        amount : 5,
        tag_id : 5
    },
    {
        post_id : 1,
        pin : 123,
        product_name : "dog",
        post_date : "12-3-63",
        product_option : "buy",
        price : 520,
        amount : 5,
        tag_id : 5
    },
    {
        post_id : 1,
        pin : 123,
        product_name : "bird",
        post_date : "12-3-63",
        product_option : "buy",
        price : 520,
        amount : 5,
        tag_id : 5
    },
    {
        post_id : 1,
        pin : 123,
        product_name : "fish",
        post_date : "12-3-63",
        product_option : "buy",
        price : 520,
        amount : 5,
        tag_id : 5
    }, 
]

const handleChange = (e) => {
    if(!tag_want.includes(e))
    {
        tag_want.push(e)
        alert(tag_want)
    }
    else {
       tag_want =  tag_want.filter(function(ele){ 
            return ele != e; 
        })
        alert(tag_want)
    }
}


export default function Matching() {
  return (
    <div className={styles.container}>
      <div>
          <div className={styles.logo} > <Image width={71} height={68} src={shopping_cart} alt="shopping_cart" /> Market</div>
          <input  placeholder='Search ?'></input>
          <button>Search</button>
      </div>
      <main className={styles.main}>
      
      {data1.map(content => (
          
       
        <div className={styles.card}>
            <div style={{color: "red", display: 'inline'}}>
                <div className={styles.logo} > <Image width={171} height={168} src={shopping_cart} alt="shopping_cart" /> {content.product_name}</div>
                 
            </div>
            
            ราคา : {content.price} / ชิ้น <br></br>
            จำนวน : {content.amount} ชิ้น
            
            </div> 

      ))}
    
      </main>
      {/* <div className={styles.next}>
        <button className={styles.nextt}>Next</button>
        <Image src="/next_icon.svg" alt="Vercel Logo" width={62} height={16} />
      </div> */}
    </div>
  )
}
