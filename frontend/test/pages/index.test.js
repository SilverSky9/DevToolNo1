import React from "react";
// Using render and screen from test-utils.js instead of
// @testing-library/react
import { screen, fireEvent, render } from '@testing-library/react';
import Enzyme, { shallow, mount } from "enzyme";
import Adapter from '@wojtekmaj/enzyme-adapter-react-17';
import Index from "../../pages/index";

Enzyme.configure({ adapter: new Adapter() });

const tag = [
  'com',
  'food',
  'it',
];

describe("HomePage", () => {

  it("should render Index", () => {
    shallow(<Index />);
  });

  it("render all tag", () => {
    const wrapper = mount(<Index allTag={tag} />);

    expect(wrapper.find('ul').children().length).toBe(3)
  });

  it("should add data to tag want", () => {
    const mockFn = jest.fn()
    const { queryByTestId } = render(
      <Index allTag={tag} onClick={mockFn} />,
    );
    fireEvent.click(queryByTestId('button0'));

    expect(queryByTestId('tagWant').childNodes.length).toBe(1);

  });

});