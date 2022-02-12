import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Matching.module.css'
import shopping_cart from '../public/shopping-cart.png' 

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
        <div>
          <button className={styles.button}>เครื่องครัว</button>
          <button className={styles.button}>เครื่องใช้</button>
          <button className={styles.button}>เครื่องเขียน</button>
          <button className={styles.button}>ต้นไม้</button>
          <button className={styles.button}>เครื่องดนตรี</button>
          <button className={styles.button}>อาหาร</button>
        </div>
        <div>
          <button className={styles.button}>เครื่องดื่ม</button>
          <button className={styles.button}>วิศวกรรมศาสตร์</button>
          <button className={styles.button}>IT</button>
          <button className={styles.button}>สถาปัตยกรรม</button>
          <button className={styles.button}>วิทยาศาสตร์</button>
        </div>
        <div>
          <button className={styles.button}>Male</button>
          <button className={styles.button}>Female</button>
          <button className={styles.button}>สัตว์เลี้ยง</button>
          <button className={styles.button}>งานอดิเรก</button>
          <button className={styles.button}>เครื่องนอน</button>
          <button className={styles.button}>ยานยนต์</button>
        </div>
      </main>
      
      <div className={styles.next}>
        <button className={styles.nextt}>Next</button>
        <Image src="/next_icon.svg" alt="Vercel Logo" width={62} height={16} />
      </div>

    </div>
  )
}
