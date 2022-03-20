import React from "react";
// Using render and screen from test-utils.js instead of
// @testing-library/react
import { fireEvent, render } from '@testing-library/react';
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

});