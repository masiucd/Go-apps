import * as React from 'react';
import { StyledMenuWrapper, StyledMenu } from './Styles.Menu';
import { graphql, useStaticQuery } from 'gatsby';
import { FixedObject } from 'gatsby-image';
import MenuItem from './MenuItem';
import Title from '../../styled/title/Title';
const PRODUCTS_QUERY = graphql`
  {
    dishes: allContentfulMenu {
      edges {
        node {
          id
          name
          ingredients
          price
          vegiterian
          img {
            fixed(width: 200, height: 200) {
              ...GatsbyContentfulFixed_tracedSVG
            }
          }
        }
      }
    }
  }
`;

interface P {
  dishes: {
    edges: {
      node: {
        id: string;
        name: string;
        ingredients: string;
        price: number;
        vegiterian: boolean;
        img: { fixed?: FixedObject | FixedObject[] | undefined };
      };
    }[];
  };
}

const Menu: React.FC<P> = () => {
  const data: P = useStaticQuery(PRODUCTS_QUERY);
  return (
    <StyledMenuWrapper>
      <Title
        mainTitle="Hungry"
        subTitle="you are?"
        flex
        ctaText="Menu"
        cta
        linkTo="/menu"
      />
      <StyledMenu>
        {data.dishes.edges.map(x => (
          <MenuItem key={x.node.id} data={x} />
        ))}
      </StyledMenu>
    </StyledMenuWrapper>
  );
};
export default Menu;
