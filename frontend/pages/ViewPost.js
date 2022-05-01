import styles from '../styles/Home.module.css'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/router'
import { Button, Modal, Row, Form, Container, Col, InputGroup, FormControl, ButtonGroup, ToggleButton, radioValue, FloatingLabel, Card } from 'react-bootstrap';
import axios from 'axios';

function ViewPost() {
  return (
    <Container className={styles.container}>
        <Col className={styles.main}>
            <div className='row '>แปรงแต่งหน้า</div>
        </Col>
        <Col className={styles.main}>
            <div className='row '>Buyer Information</div>
        </Col>
        <Col className={styles.main}>
            <Row className='row '>Buyer Information</Row>
            <Row className='row '>Buyer Information</Row>
        </Col>
    </Container>
  );
}

export default ViewPost;
