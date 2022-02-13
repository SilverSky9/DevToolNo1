import Head from 'next/head'
import styles from '../styles/Matching.module.css'
import shopping_cart from '../public/shopping-cart.png'
import { useState, useEffect } from "react";
import axios from 'axios'

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
    }

    function check() {
        console.log(tag_id)
    }


    const GetTag = async () => {
        await axios.get("http://localhost:3000/tag/getall")
            .then(res => {
                setTag(res.data)
            })
    }

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

                <hr></hr>
                <ul>
                    {tag_want.map(tag => (
                        <button className={styles.button1} key="test">{tag}</button>
                    ))}
                </ul>
            </main>
            <div className={styles.next}>
                <button className={styles.nextt}>Next</button>
            </div>
        </div>
    )
}


