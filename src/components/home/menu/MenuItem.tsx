import * as React from 'react';
import { FixedObject } from 'gatsby-image';
import { StyledMenuItem, ImgWrap, ItemBody } from './Styles.Menu';
import Img from 'gatsby-image';
import { BtnLink } from '../../styled/Buttons';

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
  const { name, price, img } = data.node;
  console.log('MENUITEM  ', data);
  return (
    <StyledMenuItem>
      <ImgWrap>
        <Img fixed={img.fixed} />
      </ImgWrap>
      <ItemBody className="body">
        <span>{name}</span>
        <span>{price}$</span>
        <BtnLink fade to="/menu">
          Menu
        </BtnLink>
      </ItemBody>
    </StyledMenuItem>
  );
};
export default MenuItem;
