import styles from '../styles/Home.module.css'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/router'
import { Button, Modal, Row, Form, Container, Col, InputGroup, FormControl, ButtonGroup, ToggleButton, radioValue, FloatingLabel, Card } from 'react-bootstrap';
import axios from 'axios';
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

var locations = [
    {
        location_id: 1,
        location_name: "RNP"
    },
    {
        location_id: 2,
        location_name: "Jinda"
    },
    {
        location_id: 3,
        location_name: "V Condo"
    }]

const Matching = ({ tag }) => {

    const radios = [
        { name: 'Buyer', value: 'buy' },
        { name: 'Seller', value: 'sell' },

    ];

    const [show, setShow] = useState(false);
    const handleClose = () => {
        setShow(false)
        setProductName('')
        setDetail('')
        setPhone('')
        setRadioValue("buy")
        setLocation('')
        setCategory('')
        setPrice('')
        setAmount('')
        setValidated(false)


    };
    const handleShow = () => setShow(true);


    const [productName, setProductName] = useState('');
    const [detail, setDetail] = useState('');
    const [phone, setPhone] = useState('');
    const [radioValue, setRadioValue] = useState("buy");
    const [location, setLocation] = useState('');
    const [category, setCategory] = useState('');
    const [price, setPrice] = useState();
    const [amount, setAmount] = useState();

    const [validated, setValidated] = useState(false);




    const [posted, setPost] = useState([])
    // const [taged, setTag] = useState([])
    const [searchVal, setSearchVal] = useState('')
    const router = useRouter()
    useEffect(() => {
        const url = router.query.tag ||= tag.map(tag => (tag.tag_id)).toString()

        async function getPost() {
            let response = await fetch("http://159.223.45.216:3010/post/geybymultitag/" + url + ',')
            response = await response.json()
            console.log(response);
            setPost(response)
        }

        getPost()
    }, [])

    const GetPostBySearch = async () => {
        const res = await fetch("http://159.223.45.216:3010/post/searchbyname/" + searchVal);
        const newPost = await res.json();

        return setPost(newPost);
    };

    const GetPostByTag = async (tag_name) => {
        const res = await fetch("http://159.223.45.216:3010/post/getbytag/" + tag_name + "/null");
        const newPost = await res.json();

        return setPost(newPost);
    };

    const addPost = async () => {
     
        var data = {
            product_name: productName,
            product_option: radioValue,
            location: location,
            phone_number: phone,
            price: parseInt(price),
            amount: parseInt(amount),
            tag_name: category

        }
        await axios.post('http://159.223.45.216:3010/post/create', data)
            .then(response => {
                console.log("add post success")
                console.log(response);
            
            })
            handleClose()
            setTimeout('alert("sucess");', 500);
        setShow(false)
    

    }

    const handleSubmit = (event) => {
        const form = event.currentTarget;
        if (form.checkValidity() === false) {
          event.preventDefault();
          event.stopPropagation();
        }
        else { 
            addPost()
            event.preventDefault();
        }
    
        setValidated(true);
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
                            // <div key={content.post_id} className={styles.card} >
                            <div key={content.post_id}>
                                <Card className={styles.card}  >
                                    <Card.Header id="productName" className="h2 text-center">{content.product_name}</Card.Header>
                                    <Card.Body>
                                        {/* <Card.Title>{content.product_name}</Card.Title> */}
                                        <Card.Text>
                                            ราคา : {content.price} / ชิ้น <br></br>
                                            จำนวน : {content.amount} ชิ้น
                                        </Card.Text>
                                    </Card.Body>
                                    <Card.Footer className="text-muted " style={{ textAlign: 'right' }} >{content.post_date.slice(0, 10)}</Card.Footer>
                                </Card>
                            </div>
                        ))}
                    </Row>


                    <Row>
                        <Button variant="primary" onClick={handleShow} className={styles.btnRight}>
                        +
                    </Button>

                    </Row>
                    
                    <Modal
                        show={show}
                        onHide={handleClose}
                        backdrop="static"
                        keyboard={false}
                        size="lg"
                        aria-labelledby="contained-modal-title-vcenter"
                        centered
                    >
                        {/* <Modal.Header closeButton>
                        <Modal.Title >Modal title</Modal.Title>
                    </Modal.Header> */}
                        <Modal.Body>
                            <Container>
                                <Form noValidate validated={validated} onSubmit={handleSubmit}>
                                    <Row>
                                        <Col>

                                            {/* {productName} */}
                                            <Form.Group md="4" controlId="validationCustom01" className="mb-3 mt-4">
                                                <FloatingLabel controlId="floatingInputGrid" label="Product Name">
                                                    <Form.Control
                                                        required
                                                        type="text"
                                                        placeholder="Product Name"
                                                        defaultValue={productName}
                                                        onChange={(e) => setProductName(e.target.value)}
                                                        data-testid="product_name"
                                                        name='productName'
                                                    />
                                                </FloatingLabel>
                                                <Form.Control.Feedback>Looks good!</Form.Control.Feedback>
                                            </Form.Group>

                                            {/* {detail} */}
                                            {/* <InputGroup className="mb-3">
                                                <InputGroup.Text id="inputGroup-sizing-default">Detail</InputGroup.Text>
                                                <FormControl
                                                    aria-label="Default"
                                                    aria-describedby="inputGroup-sizing-default"
                                                    value={detail}
                                                    onChange={(e) => setDetail(e.target.value)}
                                                />
                                            </InputGroup> */}
                                            <Form.Group md="4" controlId="validationCustom02" className="mb-3 mt-4">
                                                <FloatingLabel controlId="floatingInputGrid" label="Detail">
                                                    <Form.Control
                                                        required
                                                        type="text"
                                                        placeholder="Detail"
                                                        defaultValue={detail}
                                                        onChange={(e) => setDetail(e.target.value)}
                                                        data-testid="product_detail"
                                                        name='productDetail'
                                                    />
                                                </FloatingLabel>
                                                <Form.Control.Feedback>Looks good!</Form.Control.Feedback>
                                            </Form.Group>


                                            {/* {phone} */}

                                            <Form.Group md="4" controlId="validationCustom03" className="mb-3 mt-4">
                                                <FloatingLabel controlId="floatingInputGrid" label="Phone">
                                                    <Form.Control
                                                        required
                                                        type="number"
                                                        placeholder="Phone"
                                                        defaultValue={phone}
                                                        onChange={(e) => setPhone(e.target.value)}
                                                        data-testid="phone"
                                                        name="phone"
                                                    />
                                                </FloatingLabel>
                                                <Form.Control.Feedback>Looks good!</Form.Control.Feedback>
                                            </Form.Group>


                                            {/* {location} */}
                                            <Form.Select aria-label="Default select example" className="mb-3" onChange={(e) => setLocation(e.currentTarget.value)} required data-testid="location" name='location'>
                                                <option value="" selected disabled>Location</option>
                                                {locations?.map((item, i) => (
                                                    <option value={item.location_id} key={i}>{item.location_name}</option>
                                                ))}
                                            </Form.Select>

                                        </Col>


                                        <Col className="justify-content-md-center">
                                            {/* {radioValue} */}
                                            <ButtonGroup style={{ width: '100%' }} className="mb-3 mt-4">
                                                {radios.map((radio, idx) => (
                                                    <ToggleButton
                                                        key={idx}
                                                        id={`radio-${idx}`}
                                                        type="radio"
                                                        variant={idx % 2 ? 'outline-success' : 'outline-danger'}
                                                        name="radio"
                                                        value={radio.value}
                                                        checked={radioValue === radio.value}
                                                        onChange={(e) => setRadioValue(e.currentTarget.value)}
                                                        data-testid="toggle_button"
                                                    >
                                                        {radio.name}
                                                    </ToggleButton>
                                                ))}
                                            </ButtonGroup>

                                            {/* {price} */}
                                            <Form.Group md="4" controlId="validationCustom04" className="mb-3 mt-4">
                                                <FloatingLabel controlId="floatingInputGrid" label="Price">
                                                    <Form.Control
                                                        required
                                                        type="number"
                                                        placeholder="Price"
                                                        defaultValue={price}
                                                        onChange={(e) => setPrice(e.target.value)}
                                                        data-testid="price"
                                                        name='price'
                                                    />
                                                </FloatingLabel>
                                                <Form.Control.Feedback>Looks good!</Form.Control.Feedback>
                                            </Form.Group>

                                            {/* {amount} */}
                                            <Form.Group md="4" controlId="validationCustom05" className="mb-3 mt-4">
                                                <FloatingLabel controlId="floatingInputGrid" label="Amount">
                                                    <Form.Control
                                                        required
                                                        type="number"
                                                        placeholder="Amount"
                                                        defaultValue={amount}
                                                        onChange={(e) => setAmount(e.target.value)}
                                                        data-testid="amount"
                                                        name="amount"
                                                    />
                                                </FloatingLabel>
                                                <Form.Control.Feedback>Looks good!</Form.Control.Feedback>
                                            </Form.Group>

                                            {/* {category} */}
                                            <Form.Select aria-label="Default select example" className="mb-3" onChange={(e) => setCategory(e.currentTarget.value)} required data-testid="product_category" name='product_category'>
                                                <option value="" selected disabled>Product Category</option>
                                                {tag?.map((item, i) => (
                                                    <option value={item.tag_name} key={i}>{item.tag_name}</option>
                                                ))}
                                            </Form.Select>


                                        </Col>
                                    </Row>
                                    <Row style={{ justifyContent: 'center' }}>
                                        <Button className="mx-3" type="submit" variant="primary" style={{ width: '30%' }}
                                            data-testid="post_button"
                                        >
                                            Post
                                        </Button>



                                        <Button className="mx-3" style={{ width: '30%' }} variant="secondary" onClick={handleClose}>
                                            Cancel
                                        </Button>
                                    </Row>
                                </Form>
                            </Container>
                        </Modal.Body>

                        <Modal.Footer style={{ justifyContent: 'center' }}>



                        </Modal.Footer>
                    </Modal>
                </main>
                <div className=''>

                    {/* <button className='btn btn-primary'>create post</button> */}
                    {/* <Image src="../public/next_icon.svg" alt="Vercel Logo" width={62} height={16} /> */}

                </div>
            </div>
            <div className='col-2 bg-light '>
                <div className='row mt-4 position-fixed'>
                    <div style={{ color: 'rgb(75, 75, 75)' }}>
                        <ul>
                            {tag?.map((item, i) => (
                                <div key={i} className={`tag ${styles.tag}  bg-dark text-light`} onClick={() => {
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
Matching.getInitialProps = async () => {
    // const router = useRouter()
    const tag = await fetch("http://159.223.45.216:3010/tag/getall")
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
