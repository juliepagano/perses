// Copyright 2023 The Perses Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import CodeIcon from 'mdi-material-ui/CodeBraces';
import { InfoTooltip } from '@perses-dev/components';
import { TOOLTIP_TEXT } from '../../constants';
import { ToolbarIconButton } from '../ToolbarIconButton';
import { useEditJsonDialog } from '../../context';

export const EditJsonButton = () => {
  const { openEditJsonDialog } = useEditJsonDialog();
  return (
    <InfoTooltip description={TOOLTIP_TEXT.editJson}>
      <ToolbarIconButton variant="outlined" onClick={() => openEditJsonDialog()}>
        <CodeIcon />
      </ToolbarIconButton>
    </InfoTooltip>
  );
};
