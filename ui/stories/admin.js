import React from 'react'
import {storiesOf, action, linkTo} from '@kadira/storybook'
import Center from './components/Center'

import MultiSelectDropdown from 'shared/components/MultiSelectDropdown'
import Tooltip from 'shared/components/Tooltip'
import TooltipDelete from 'src/admin/components/TooltipDelete'

storiesOf('MultiSelectDropdown', module)
  .add('Select Roles', () => (
    <Center>
      <MultiSelectDropdown
        items={[
          'Admin',
          'User',
          'Chrono Giraffe',
          'Prophet',
          'Susford',
        ]}
        selectedItems={[
          'User',
          'Chrono Giraffe',
        ]}
        label={'Select Roles'}
        onApply={action('onApply')}
      />
    </Center>
  ))

storiesOf('Tooltip', module)
  .add('Delete', () => (
    <Center>
      <Tooltip tip={`Are you sure? TrashIcon`}>
        <div className="btn btn-info btn-sm">
          Delete
        </div>
      </Tooltip>
    </Center>
  ))

storiesOf('TooltipDelete', module)
  .add('Delete', () => (
    <Center>
      <TooltipDelete onDelete={action('onDelete')} />
    </Center>
  ))