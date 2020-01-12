import * as React from 'react';
import { FixedObject } from 'gatsby-image';
import Img from 'gatsby-image';
import {
  StyledMenuItem,
  StyledMenuItemHead,
  StyledMenuItemFooter,
} from './Menu.Styles';
interface P {
  data: {
    node: {
      id: string;
      name: string;
      ingredients: string;
      price: number;
      vegiterian: boolean;
      img: { fixed?: FixedObject | FixedObject[] | undefined };
    };
  };
}

const MenuItem: React.FC<P> = ({ data }) => {
  const { vegiterian, name, price, ingredients, img } = data.node;
  return (
    <StyledMenuItem>
      <StyledMenuItemHead>
        <Img fixed={img.fixed} />
      </StyledMenuItemHead>
      <StyledMenuItemFooter vegetarian={vegiterian}>
        <span>title {name} </span>
        <span>price: {price}$ </span>
        <span> Vegetarian: {vegiterian ? ' Oui' : ' Noo'}</span>
      </StyledMenuItemFooter>
    </StyledMenuItem>
  );
};
export default MenuItem;
