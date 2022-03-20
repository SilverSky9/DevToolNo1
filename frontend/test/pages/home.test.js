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
describe("HomePage", () => {

    it("should render Home", () => {
        shallow(<Home tag={tag} />);
    });

    // it("render all tag", () => {
    //     const wrapper = mount(<Index allTag={tag} />);

    //     expect(wrapper.find('ul').children().length).toBe(3)
    // });


});