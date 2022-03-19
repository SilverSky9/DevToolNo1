import React from "react";
// Using render and screen from test-utils.js instead of
// @testing-library/react
import Enzyme, { shallow, mount } from "enzyme";
import Adapter from '@wojtekmaj/enzyme-adapter-react-17';
import Index from "../../pages/index";

Enzyme.configure({ adapter: new Adapter() });

const tag = [
  'com',
  'food',
  'it'
];

describe("HomePage", () => {
  it("should render Index", () => {
    shallow(<Index />);
  });

  it("render three tag", () => {
    const wrapper = mount(<Index allTag={tag} />);
    // console.log(wrapper.find("#button").length);
    expect(wrapper.find("button")[0]).toBe(4)
  });

  it("should render tag that you selected", () => {
    const wrapper = mount(<Index allTag={tag} />);

    console.log(wrapper.find("#button"));
    expect(wrapper.find("#button0").value)
  });

});