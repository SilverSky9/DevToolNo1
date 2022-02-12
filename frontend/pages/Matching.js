import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Matching.module.css'
import shopping_cart from '../public/shopping-cart.png' 

const tag_want = []
var data1 = [
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
const isActive = false
const handleChange = (e) => {
    if(!tag_want.includes(e))
    {
        tag_want.push(e)
        alert(tag_want)
        isActive = true
        alert(isActive)
    }
    else {
        isActive = false
        alert(isActive)
       tag_want =  tag_want.filter(function(ele){ 
            return ele != e; 
        })
        alert(tag_want)
    }
}
// const isActive = (r) => {
//     // alert(r)

//     if (tag_want.includes(r)) {
//         alert("incluse")
//         return 1
//     }
//     else return 0
// }




export default function Matching() {
  return (
    <div className={styles.container}>
      <div>
          <h1 className={styles.title}><a style={{ color: '#F49A35' }}>IN</a> <a style={{ color: '#197DFF' }}>this</a></h1>
          <div className={styles.logo} > <Image width={71} height={68} src={shopping_cart} alt="shopping_cart" /> Market</div>
          <h1 className={styles.title}>
            <a style={{ color: '#197DFF' }}>What</a> do you <a style={{ color: '#F49A35' }}>want</a> <a style={{ color: '#197DFF' }}>?</a>
          </h1>
      </div>
      <main className={styles.main}>
      <ul>
      {data1.map(tag => (
          
      
        // <button  onClick={() => handleChange(tag.tag_id)} className={ isActive(tag.tag_id) ? styles.button1 : styles.button2}>{tag.tag_name}</button> 
        <button  onClick={() => handleChange(tag.tag_id)} className={isActive ? styles.button1 : styles.button2}>{tag.tag_name}</button> 

      ))}
    </ul>
      </main>
      <div className={styles.next}>
        <button className={styles.nextt}>Next</button>
        <Image src="/next_icon.svg" alt="Vercel Logo" width={62} height={16} />
      </div>
    </div>
  )
}
