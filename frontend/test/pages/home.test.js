import React from "react";
// Using render and screen from test-utils.js instead of
// @testing-library/react
import { fireEvent, render, screen } from '@testing-library/react';
import Enzyme, { shallow, mount } from "enzyme";
import Adapter from '@wojtekmaj/enzyme-adapter-react-17';
import Home from "../../pages/Home";

Enzyme.configure({ adapter: new Adapter() });


const tag = [
    'com',
    'food',
    'it',
  

  
];

const posted = [
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
        post_id: 2,
        pin: 123,
        product_name: "dog",
        post_date: "12-3-63",
        product_option: "buy",
        price: 520,
        amount: 5,
        tag_id: 5
    },
]

describe("HomePage", () => {

    it("should render Home", () => {
        shallow(<Home tag={tag} />);
    });

    it("render all tag", () => {
        const wrapper = shallow(<Home posted={posted} tag={tag} />);
        expect(wrapper.find('ul').children().length).toBe(3)
    });

    it("render all posts", () => {
        const wrapper = shallow(<Home posted={posted} tag={tag} />);
        expect(wrapper.find('ul').children().length).toBe(3)
    });

    it("test input and toggle button is defined in the modal.", () => {
        // render(<Home />)
        
        expect(screen.queryByTestId("product_name")).toBeDefined();
        expect(screen.queryByTestId("product_detail")).toBeDefined();
        expect(screen.queryByTestId("phone")).toBeDefined();
        expect(screen.queryByTestId("location")).toBeDefined();
        expect(screen.queryByTestId("toggle_button")).toBeDefined();
        expect(screen.queryByTestId("price")).toBeDefined();
        expect(screen.queryByTestId("amount")).toBeDefined();
        expect(screen.queryByTestId("product_category")).toBeDefined();
    })

    it("test input received the expected value.", () => {
        const testState = { 
            product_name: 'ยางลบ', 
            product_detail: 'ยางลบจ้า',
            phone: '0123456789',
            location: 'RNP',
            // toggle_button: 'Buyer',
            price: '20',
            amount: '0',
            product_category: 'เครื่องเขียน'
        };

        const wrapper = shallow(<Home posted={posted} tag={tag} />);

        wrapper.find('input').at(0).simulate('change', { target: { name: 'productName', value: 'ยางลบ' } });
        wrapper.find('input').at(0).simulate('change', { target: { name: 'productDetail', value: 'ยางลบจ้า' } });
        wrapper.find('input').at(0).simulate('change', { target: { name: 'phone', value: '0123456789' } });
        wrapper.find('input').at(0).simulate('change', { target: { name: 'location', value: 'RNP' } });
        wrapper.find('input').at(0).simulate('change', { target: { name: 'price', value: '20' } });
        wrapper.find('input').at(0).simulate('change', { target: { name: 'amount', value: '0' } });
        wrapper.find('input').at(0).simulate('change', { target: { name: 'product_category', value: 'เครื่องเขียน' } });

        expect(testState.product_name).toEqual('ยางลบ');
        expect(testState.product_detail).toEqual('ยางลบจ้า');
        expect(testState.phone).toEqual('0123456789');
        expect(testState.location).toEqual('RNP');
        expect(testState.price).toEqual('20');
        expect(testState.amount).toEqual('0');
        expect(testState.product_category).toEqual('เครื่องเขียน');
    })

});