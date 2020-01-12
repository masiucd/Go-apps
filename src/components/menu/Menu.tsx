import * as React from 'react';
import { StyledMenu, Grid } from './Menu.Styles';
import MenuItem from './MenuItem';
import { graphql, useStaticQuery } from 'gatsby';
import { FixedObject } from 'gatsby-image';
import Title from '../styled/title/Title';

const MENU_QUERY = graphql`
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
            fixed(width: 300, height: 300) {
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
  const data: P = useStaticQuery(MENU_QUERY);
  return (
    <StyledMenu>
      <Title mainTitle="We Love food" subTitle="and you ?" flex />
      <Grid>
        {data.dishes.edges.map(x => (
          <MenuItem key={x.node.id} data={x} />
        ))}
      </Grid>
    </StyledMenu>
  );
};
export default Menu;
