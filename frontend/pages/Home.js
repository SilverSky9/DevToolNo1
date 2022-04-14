import styles from '../styles/Home.module.css'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/router'
import { Button, Modal, Row } from 'react-bootstrap';
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








const Matching = ({ tag }) => {
    const [show, setShow] = useState(false);

    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);


    const [posted, setPost] = useState([])
    // const [taged, setTag] = useState([])
    const [searchVal, setSearchVal] = useState('')
    const router = useRouter()
    useEffect(() => {
        const url = router.query.tag ||= tag.map(tag => (tag.tag_id)).toString()

        async function getPost() {
            let response = await fetch("http://localhost:3000/post/geybymultitag/" + url + ',')
            response = await response.json()
            console.log(response);
            setPost(response)
        }

        getPost()
    }, [])

    const GetPostBySearch = async () => {
        const res = await fetch("http://localhost:3000/post/searchbyname/" + searchVal);
        const newPost = await res.json();

        return setPost(newPost);
    };

    const GetPostByTag = async (tag_name) => {
        const res = await fetch("http://localhost:3000/post/getbytag/" + tag_name + "/null");
        const newPost = await res.json();

        return setPost(newPost);
    };

    return (
        <div className='row '>
            {/* {router.query.tag.map(tag => <h1>{tag}<h1 /> )} */}
            <div className='col-10 px-5'>
                <div className='row mt-4'>
                    <div className='col-10 '>
                        <input className='w-100 mt-2 p-2 form-control' onChange={e => setSearchVal(e.target.value)}></input>
                    </div>
                    <div className='col-2'>
                        <button className='btn btn-primary mt-2 w-100' id="search" onClick={() => GetPostBySearch()}>Search</button>
                    </div>
                </div>
                <main style={{ textAlign: 'center' }}>
                    <Row style={{ flex: 1, justifyContent: 'center' }}>
                        {posted?.map(content => (
                            <div key={content.post_id} className={styles.card} >
                                <div id="productName" style={{ color: '#197DFF', fontSize: '50px', textAlign: 'center' }} >
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
                    </Row>
                    <Button variant="primary" onClick={handleShow} className={styles.btnRight}>
                        Create Post
                    </Button>

                    <Modal
                        show={show}
                        onHide={handleClose}
                        backdrop="static"
                        keyboard={false}
                        size="lg"
                        aria-labelledby="contained-modal-title-vcenter"
                        centered
                    >
                        <Modal.Header closeButton>
                            <Modal.Title >Modal title</Modal.Title>
                        </Modal.Header>
                        <Modal.Body>
                            I will not close if you click outside me. Don't even try to press
                            escape key.
                        </Modal.Body>
                        <Modal.Footer>
                            <Button variant="secondary" onClick={handleClose}>
                                Close
                            </Button>
                            <Button variant="primary">Understood</Button>
                        </Modal.Footer>
                    </Modal>
                </main>
                <div className=''>

                    {/* <button className='btn btn-primary'>create post</button> */}
                    {/* <Image src="../public/next_icon.svg" alt="Vercel Logo" width={62} height={16} /> */}

                </div>
            </div>
            <div className='col-2 bg-light '>
                <div className='row mt-4 position-fixed '>
                    <div style={{ color: 'rgb(75, 75, 75)' }}>
                        <ul>
                            {tag?.map((item, i) => (
                                <div key={i} className={`tag ${styles.tag}`} onClick={() => {
                                    GetPostByTag(item.tag_id)
                                }}  >
                                    <div > {item.tag_name}
                                    </div>
                                </div>
                            ))}
                        </ul>
                    </div>
                </div>
            </div>






        </div>


    )
}

Matching.getInitialProps = async (ctx) => {
    // const router = useRouter()
    const tag = await fetch("http://localhost:3000/tag/getall")
    const allTag = await tag.json()

    // const url = context.query.tag ||= allTag.map(tag => (tag.tag_id)).toString()

    // const selectedTag = await fetch("http://34.126.190.231:3000/post/geybymultitag/" + url + ',')
    // const selectTag = await selectedTag.json()

    return {

        // post: selectTag,
        tag: allTag
        // will be passed to the page component as props
    }
}

export default Matching
