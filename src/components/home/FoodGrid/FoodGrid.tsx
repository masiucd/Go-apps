import * as React from 'react';
import { StyledFoodGrid, GridItem } from './Styles.foodgrid';

interface P {}

const FoodGrid: React.FC<P> = () => {
  return (
    <StyledFoodGrid>
      <GridItem className="main">
        {' '}
        <img
          src="https://images.unsplash.com/photo-1484723091739-30a097e8f929?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=300&q=60"
          alt=""
        />{' '}
      </GridItem>
      <GridItem>
        {' '}
        <img
          src="https://images.unsplash.com/photo-1540189549336-e6e99c3679fe?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=300&q=60"
          alt=""
        />{' '}
      </GridItem>
      <GridItem>
        <img
          src="https://images.unsplash.com/photo-1511690656952-34342bb7c2f2?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=300&q=60"
          alt=""
        />
      </GridItem>
    </StyledFoodGrid>
  );
};
export default FoodGrid;
