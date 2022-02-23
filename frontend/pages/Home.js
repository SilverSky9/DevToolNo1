// import Head from 'next/head'
import styles from '../styles/Home.module.css'
import shopping_cart from '../public/shopping-cart.png'
import axios from 'axios'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/router'

const tag_want = []
var data1 = [
    {
        post_id: 1,
        pin: 123,
        product_name: "cat",
        post_date: "12-3-63",
        product_option: "buy",
        price: 520,
        amount: 5,
        tag_id: 5
    },
    {
        post_id: 2,
        pin: 123,
        product_name: "dog",
        post_date: "12-3-63",
        product_option: "buy",
        price: 520,
        amount: 5,
        tag_id: 5
    },
    {
        post_id: 3,
        pin: 123,
        product_name: "bird",
        post_date: "12-3-63",
        product_option: "buy",
        price: 520,
        amount: 5,
        tag_id: 5
    },
    {
        post_id: 4,
        pin: 123,
        product_name: "fish",
        post_date: "12-3-63",
        product_option: "buy",
        price: 520,
        amount: 5,
        tag_id: 5
    },
]


var data2 = [
    {
        tag_id: 1,
        tag_name: "เครื่องใช้"
    },
    {
        tag_id: 2,
        tag_name: "เครื่องเขียน"
    },
    {
        tag_id: 3,
        tag_name: "เครื่องดนตรี"
    },
    {
        tag_id: 4,
        tag_name: "เครื่องดื่ม"
    },
    {
        tag_id: 5,
        tag_name: "วิศวกรรมศาสตร์"
    },
    {
        tag_id: 6,
        tag_name: "IT"
    },
    {
        tag_id: 7,
        tag_name: "วิทยาศาสตร์"
    },
    {
        tag_id: 8,
        tag_name: "สัตว์เลี้ยง"
    }, {
        tag_id: 9,
        tag_name: "งานอดิเรก"
    }, {
        tag_id: 10,
        tag_name: "เครื่องนอน"
    }, {
        tag_id: 11,
        tag_name: "ยานยนต์"
    },
]



export default function Matching() {
    const [post, setPost] = useState([])
    const [tag, setTag] = useState([])
    const [searchVal, setSearchVal] = useState('')
    const router = useRouter()

    useEffect(() => {
        GetPost()
        GetTag()
        return () => {

        }
    }, [])

    const GetPost = async () => {
        await axios.get("http://34.126.190.231:3000/post/all")
            .then(res => {

                setPost(res.data)
            })
    }

    const GetTag = async () => {
        await axios.get("http://34.126.190.231:3000/tag/getall")
            .then(res => {
                setTag(res.data)
            })
    }

    const GetPostBySearch = async () => {
        await axios.get("http://34.126.190.231:3000/post/searchbyname/" + searchVal)
            .then(res => {

                setPost(res.data)
            })
    }

    const GetPostByTag = async (tag_name) => {
        await axios.get("http://34.126.190.231:3000/post/getbytag/" + tag_name + "/null")
            .then(res => {
                setPost(res.data)
            })
    }

    return (
        <div className='row '>

            {console.log(router.query.tag)}
            <div className='col-10 px-5'>
                <div className='row mt-4'>
                    <div className='col-10 '>
                        <input className='w-100 mt-2 p-2 form-control' onChange={e => setSearchVal(e.target.value)}></input>
                    </div>
                    <div className='col-2'>
                        <button className='btn btn-primary mt-2 w-100' onClick={() => GetPostBySearch()}>Search</button>
                    </div>
                </div>
                <img src='../public/shopping-cart.png' />





                <main className={styles.main}>


                    {post.map(content => (
                        <div key={content.post_id} className={styles.card} >



                            <div style={{ color: '#197DFF', fontSize: '50px', textAlign: 'center' }} >
                                {/* <div className={styles.logo} > <Image width={171} height={168} src={shopping_cart} alt="shopping_cart" /> {content.product_name}</div> */}
                                {content.product_name} <br></br>
                            </div>
                            <div >

                                ราคา : {content.price} / ชิ้น <br></br>
                                จำนวน : {content.amount} ชิ้น
                            </div>
                            {/* <div className={styles.logo} > <Image width={171} height={168} src={shopping_cart} alt="shopping_cart" /> {content.product_name}</div> */}


                        </div>


                    ))}

                </main>
                <div className=''>

                    <button className='btn btn-primary'>Next</button>
                    {/* <Image src="../public/next_icon.svg" alt="Vercel Logo" width={62} height={16} /> */}

                </div>
            </div>
            <div className='col-2 bg-light'>
                <div className='row mt-4'>
                    <div style={{ color: 'rgb(75, 75, 75)' }}>
                        {tag.map(tag => (
                            <div key={tag.tag_id} className={styles.tag} onClick={() => {
                                GetPostByTag(tag.tag_id)
                            }}  >
                                <div > {tag.tag_name}
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            </div>
        </div>
    )
}
