import styles from '../styles/Matching.module.css'
import home_styles from '../styles/Home.module.css'

import { useState, useEffect } from "react";
import Link from 'next/link'
import { Button } from 'react-bootstrap';



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




const Matching = ({ allTag = [] }) => {
    const [tag_want, setTag_want] = useState([]);
    const [tag_id, setTag_id] = useState([]);

    const [tag_want_filter, set_tag_want_filter] = useState('')

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
    }

    // const GetPostFromTag = async () => {
    //     await axios.get("http://34.126.190.231:3000/post/geybymultitag/" + tag_want_filter)
    //         .then(res => {
    //             setPost(res.data)
    //             setToHome(true)
    //         })
    // }

    return (

        <div className={styles.container}>

            <div data-testid="custom-element">
                <h1 className={styles.title}><a style={{ color: '#F49A35' }}>IN</a> <a style={{ color: '#197DFF' }}>this</a></h1>
                <h1 className={styles.title}>
                    <a style={{ color: '#197DFF' }}>What</a> do you <a style={{ color: '#F49A35' }}>want</a> <a style={{ color: '#197DFF' }}>?</a>
                </h1>
            </div>
            <main className={styles.main}>

                <ul>
                    {allTag?.map((tag, index) => (
                        <Button key={index} id={`button `} data-testid={`${'button' + index}`} onClick={() => handleChange(tag.tag_name, tag.tag_id)} className={styles.button}>{tag.tag_name}</Button>

                    ))}
                </ul>

                <hr>

                </hr>
                <ul data-testid={'tagWant'}>
                    {tag_want?.map((tag, i) => (
                        <span className={styles.button1} data-testid={`${'span' + i}`} key={i}>{tag} </span>
                    ))}
                </ul>

            </main>
            {tag_id.length == 0 ?
                null
                :
                <Link href={{
                    pathname: '/Home',
                    query: { tag: tag_id },
                }}>
                    <div className={styles.next}>
                        <button className={styles.nextt} id="next-button"
                        // onClick={() => GetPostFromTag()}
                        >Next</button>
                    </div></Link>
            }


        </div>


    )
}



Matching.getInitialProps = async () => {
    const res = await fetch('http://159.223.45.216:3010/tag/getall')
        // const res = await fetch('http://localhost:3000/tag/getall')
       
    const tag = await res.json()

    // console.log(tag);
    return {
        allTag: tag 
    }
}

// Matching.getInitialProps = async () => {
//         const res = await fetch('http://localhost:3000/tag/getall')
       
//     const tag = await res.json()

//     console.log(tag);
//     return {
//         allTag: tag 
//     }
// }

export default Matching
