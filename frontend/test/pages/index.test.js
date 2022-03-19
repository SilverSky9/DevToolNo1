import React from "react";
// Using render and screen from test-utils.js instead of
// @testing-library/react
import { render, screen } from '@testing-library/react';
import Enzyme, { shallow, mount, to, have } from "enzyme";
import Adapter from '@wojtekmaj/enzyme-adapter-react-17';
import Index from "../../pages/index";

Enzyme.configure({ adapter: new Adapter() });

const tag = [
  'com',
  'food',
  'it',
];

describe("HomePage", () => {
  const wrapper = mount(<Index allTag={tag} />);

  it("should render Index", () => {
    shallow(<Index />);
  });

  it("render all tag", () => {
    expect(wrapper.find('ul').children().length).toBe(3)
  });

  it("should call function handleChange", () => {
    console.log(wrapper.find("#button"));
    expect(wrapper.getByTestId("#button0").value)
  });

});