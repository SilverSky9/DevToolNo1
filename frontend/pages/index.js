import Head from 'next/head'
import styles from '../styles/Matching.module.css'
import home_styles from '../styles/Home.module.css'
import shopping_cart from '../public/shopping-cart.png'

import { useState, useEffect } from "react";
import axios from 'axios'
import Link from 'next/link'


var data1 = [
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

    const [tag_want, setTag_want] = useState([]);
    const [tag_id, setTag_id] = useState([]);
    const [tag, setTag] = useState([])
    const [toHome, setToHome] = useState(false)
    const [post, setPost] = useState([{
        post_id: 1,
        pin: 123,
        product_name: "cat",
        post_date: "12-3-63",
        product_option: "buy",
        price: 520,
        amount: 5,
        tag_id: 5
    }])
    const [tag_want_filter, set_tag_want_filter] = useState('')

    useEffect(() => {
        GetTag()

        return () => {

        }
    }, [])

    function handleChange(e, id) {

        if (!tag_want.includes(e)) {
            setTag_want(oray => [...oray, e])
            setTag_id(ids => [...ids, id])
        }

        else {
            setTag_want(
                tag_want.filter((ele) =>
                    ele != e
                ))

            setTag_id(
                tag_id.filter((ids) =>
                    ids != id
                ))
        }
        FilterTagWantToAxios()

    }

    const FilterTagWantToAxios = async () => {
        var temp = ''
        tag_id.forEach(x => {
            temp += x + ','
        })
        set_tag_want_filter(temp)
        // console.log(tag_want_filter)
    }

    const GetPostFromTag = async () => {


        await axios.get("http://34.126.190.231:3000/post/geybymultitag/" + tag_want_filter)
            .then(res => {
                setPost(res.data)
                setToHome(true)
            })
    }

    function check() {
        console.log(tag_id)
    }


    const GetTag = async () => {
        await axios.get("http://34.126.190.231:3000/tag/getall")
            .then(res => {
                setTag(res.data)

            })
    }

    const HomeComp = () => {
        return (

            <main className={home_styles.main}>


                {post.map(content => (
                    <div key={content.post_id} className={home_styles.card} >



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



        )
    }

    const MatchingComp = () => {
        return (

            <div className={styles.container}>

                <div>
                    <h1 className={styles.title}><a style={{ color: '#F49A35' }}>IN</a> <a style={{ color: '#197DFF' }}>this</a></h1>
                    <h1 className={styles.title}>
                        <a style={{ color: '#197DFF' }}>What</a> do you <a style={{ color: '#F49A35' }}>want</a> <a style={{ color: '#197DFF' }}>?</a>
                    </h1>
                </div>
                <main className={styles.main}>

                    <ul>
                        {tag.map(tag => (
                            <button key={tag.tag_id} onClick={() => handleChange(tag.tag_name, tag.tag_id)} className={styles.button}>{tag.tag_name}</button>

                        ))}
                    </ul>

                    <hr>

                    </hr>
                    <ul>
                        {tag_want.map((tag, i) => (
                            <span className={styles.button1} key={i}>{tag}</span>
                        ))}
                    </ul>

                </main>
                <Link href={{
                    pathname: '/Home',
                    query: { tag: tag_id },
                }}>
                    <div className={styles.next}>
                        <button className={styles.nextt}
                        // onClick={() => GetPostFromTag()}
                        >Next</button>
                    </div></Link>



            </div>


        )
    }
    if (toHome) {
        return HomeComp()
    }
    else {
        return MatchingComp()
    }



}


