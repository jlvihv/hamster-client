/**
 *  Introduces component library styles on demand.
 * https://github.com/anncwb/vite-plugin-style-import
 */
import { createStyleImportPlugin, AndDesignVueResolve } from 'vite-plugin-style-import';

export function configStyleImportPlugin(_isBuild: boolean) {
  const styleImportPlugin = createStyleImportPlugin({
    resolves: [AndDesignVueResolve()],
    libs: [
      {
        libraryName: 'ant-design-vue',
        esModule: true,
        resolveStyle: (name) => {
          const ignoreList = [
            'anchor-link',
            'sub-menu',
            'menu-item',
            'menu-divider',
            'menu-item-group',
            'breadcrumb-item',
            'breadcrumb-separator',
            'form-item',
            'step',
            'select-option',
            'select-opt-group',
            'card-grid',
            'card-meta',
            'collapse-panel',
            'descriptions-item',
            'list-item',
            'list-item-meta',
            'table-column',
            'table-column-group',
            'tab-pane',
            'tab-content',
            'timeline-item',
            'tree-node',
            'skeleton-input',
            'skeleton-avatar',
            'skeleton-title',
            'skeleton-paragraph',
            'skeleton-image',
            'skeleton-button',
          ];

          // These child-component have no styles,
          // we need import's its parents
          const replaceList = {
            'typography-text': 'typography',
            'typography-title': 'typography',
            'typography-paragraph': 'typography',
            'typography-link': 'typography',
            'dropdown-button': 'dropdown',
            'input-password': 'input',
            'input-search': 'input',
            'input-group': 'input',
            textarea: 'input',
            'radio-group': 'radio',
            'checkbox-group': 'checkbox',
            'layout-sider': 'layout',
            'layout-content': 'layout',
            'layout-footer': 'layout',
            'layout-header': 'layout',
            'month-picker': 'date-picker',
            'range-picker': 'date-picker',
            'upload-dragger': 'upload',
          };

          return ignoreList.includes(name)
            ? ''
            : replaceList.hasOwnProperty(name)
            ? `ant-design-vue/es/${replaceList[name]}/style/index`
            : `ant-design-vue/es/${name}/style/index`;
        },
      },
    ],
  });

  return styleImportPlugin;
}
