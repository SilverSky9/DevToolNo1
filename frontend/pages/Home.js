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
        post_id : 2,
        pin : 123,
        product_name : "dog",
        post_date : "12-3-63",
        product_option : "buy",
        price : 520,
        amount : 5,
        tag_id : 5
    },
    {
        post_id : 3,
        pin : 123,
        product_name : "bird",
        post_date : "12-3-63",
        product_option : "buy",
        price : 520,
        amount : 5,
        tag_id : 5
    },
    {
        post_id : 4,
        pin : 123,
        product_name : "fish",
        post_date : "12-3-63",
        product_option : "buy",
        price : 520,
        amount : 5,
        tag_id : 5
    }, 
]


var data2 = [
    {
        tag_id : 1,
        tag_name : "เครื่องใช้"
    },
    {
        tag_id : 2,
        tag_name : "เครื่องเขียน"
    },
    {
        tag_id : 3,
        tag_name : "เครื่องดนตรี"
    },
    {
        tag_id : 4,
        tag_name : "เครื่องดื่ม"
    },
    {
        tag_id : 5,
        tag_name : "วิศวกรรมศาสตร์"
    },
    {
        tag_id : 6,
        tag_name : "IT"
    },
    {
        tag_id : 7,
        tag_name : "วิทยาศาสตร์"
    },
    {
        tag_id : 8,
        tag_name : "สัตว์เลี้ยง"
    },{
        tag_id : 9,
        tag_name : "งานอดิเรก"
    },{
        tag_id : 10,
        tag_name : "เครื่องนอน"
    },{
        tag_id : 11,
        tag_name : "ยานยนต์"
    },
]



export default function Matching() {
  return (
    <div className={styles.container}>
      <div>
          <div className={styles.logo} > <Image width={71} height={68} src={shopping_cart} alt="shopping_cart" /> Mar<span style={{ color: '#197DFF' }}>ket</span></div>
          <input className={styles.search} ></input>
          <button className={styles.btn}>Search</button>
      </div>
      <div style={{color: 'rgb(75, 75, 75)'}}>
         {data2.map(tag => (
        <div key={tag.tag_id} className={styles.tag}  >
            <div > {tag.tag_name}
            </div>
            </div> 
      ))} 
      </div>

    
      

      <main className={styles.main}>

      
      {data1.map(content => (
        <div key={content.post_id} className={styles.card} >
            <div style={{color: "red", display: 'inline'}}>
                {/* <div className={styles.logo} > <Image width={171} height={168} src={shopping_cart} alt="shopping_cart" /> {content.product_name}</div> */}
                 
            </div>
            {content.product_name} <br></br>
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
