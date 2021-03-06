<template>
	<div class="ability">
		<div v-if="!update" class="row mx-0">
			<span class="col-3">
				{{getHeaderName(ability.header)}} {{ ability.name }} {{ability.star | star}}<br />
				<span class="vo">{{ ability.title }}</span>
			</span>
			<span class="col-8">
				{{ ability.description }}<br />
				<span class="vo">{{ vo.description }}</span>
			</span>
			<div class="col-1 px-0">
				<div class="float-right">
					<button class="btn-success mb-1" @click="startUpdate()">Update</button>
					<button class="btn-danger" @click="$emit('remove')">Delete</button>
				</div>
			</div>
		</div>

		<div v-if="update" class="row">
			<v-autocomplete
				v-if="newAbility"
				:items="items"
				:get-label="getLabel"
				@update-items="updateItems"
				:component-item="template"
				:auto-select-one-item="false"
				@item-selected="selectedItem"
				@input="inputItem"
				placeholder="English Name"
				class="col-2 pr-0"
			></v-autocomplete>
			<input
				v-model="ability.name"
				class="col-3"
				:class="{ 'ml-3': !newAbility }"
				placeholder="Translated Name"
			/>
			<div class="form-check-inline  ml-3 col-3">
				<label>Header</label>
				<select v-model.number="ability.header">
					<option :value=null>None</option>
					<option v-for="a in itemAbilities" :key="a.id" :value="a.id">{{ a.title }}</option>
				</select>
			</div>
			<!-- <div class="form-check-inline  col-2">
				<label>Type</label>
				<select v-model.number="ability.type">
					<option value="0">None</option>
					<option value="1">Magic Ability</option>
					<option value="2">Battle Plan</option>
					<option value="3">Attack Type</option>
				</select>
			</div> -->
			<div class="form-check-inline col-3">
				<label>Star</label>
				<select v-model.number="ability.star">
					<option :value=null>None</option>
					<option value="1">*Attack</option>
					<option value="2">*Action</option>
					<option value="3">*Action or *Attack</option>
				</select>
			</div>
			<div class="col-11">
				<TextArea
					v-model="ability.description"
					:ref_id="ability_id"
					:abilities="abilitiesList"
					class="w-100"
					:rows="3"
					placeholder="Translated ability description"
				/>
			</div>
			<div class="col-1">
				<div class="float-right">
					<button v-if="ability.id && !newAbility" class="btn-success" @click="save(ability)">Update</button>
					<button v-if="ability.id && !newAbility" class="my-1" @click="cancelUpdate()">Cancel</button>
					<button v-if="newAbility" @click="save(ability)">Add</button>
				</div>
			</div>

			<div v-if="ability.id" class="col-12 vo">{{ vo.name }}: {{ vo.description }}</div>
		</div>
		<hr v-if="!newAbility" />
	</div>
</template>

<script>
import ItemTemplate from './ItemTemplate.vue'
import TextArea from './textarea.vue'

export default {
	name: 'Ability',
	props: ['abilitiesList', 'ability_id', 'ability_header', 'ability_star', 'itemAbilities'],
	components: { TextArea },
	watch: {},
	created: function() {
		if (!this.ability_id) {
			this.update = true
			this.newAbility = true
		}
		this.get(this.ability_id)
	},
	data() {
		return {
			ability: {
				type: 0,
				star: 0,
			},
			vo: {},
			template: ItemTemplate,
			items: [],
			update: false,
			beforeEdit: {},
			newAbility: false,
		}
	},
	methods: {
		startUpdate: function() {
			this.update = true
			this.beforeEdit = JSON.parse(JSON.stringify(this.ability))
		},
		cancelUpdate: function() {
			this.update = false
			this.ability = JSON.parse(JSON.stringify(this.beforeEdit))
		},
		get: function(id) {
			if (id == null) {
				return
			}
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT + `/abilities/${id}?lang=US`)
				.then(function(res) {
					console.debug(res)
					this.vo = res.data
				})
				.catch(function(err) {
					console.error(err)
				})
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT + `/abilities/${id}?lang=${this.$language}`)
				.then(function(res) {
					this.ability = res.data
					this.ability.header = this.ability_header
					this.ability.star = this.ability_star
				})
				.catch(function(err) {
					console.error(err)
				})
		},
		save: function(ability) {
			if (ability.id == null) {
				ability.id = 0
			}

			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT + `/abilities/${ability.id}?lang=${this.$language}`, ability)
				.then(function(res) {
					console.debug(res)
					if (res.status === 201) {
						ability.id = res.data
					}
					if (!this.newAbility) {
						this.$emit('add', ability, false)
						this.update = false
					} else {
						this.$emit('add', ability, true)
						this.ability = {}
					}
					this.$emit('update')
				})
				.catch(function(err) {
					console.error(err)
				})
		},

		// Handle Autocomplete
		getLabel: function(item) {
			if (!item) {
				return
			}
			return item.title
		},
		getHeaderName: function(val) {
			var res = this.itemAbilities.find(function(item){return item.id === val});
			if (res != null) {
				return '['+res.name+']'
			}
			return ''
		},
		updateItems(text) {
			this.ability.title = text
			this.ability.id = null
			this.items = this.abilitiesList
				.filter(item => item.title != null)
				.filter(item => item.title.toLowerCase().startsWith(text.toLowerCase()))
		},
		selectedItem(item) {
			this.get(item.id)
		},
		inputItem(item) {
			if (item === null) {
				this.ability = { type: 0, star: 0 }
			}
		},
	},
	filters: {
		star: function(value) {
			if (value === 1) return '(*Attack)'
			if (value === 2) return '(*Action)'
			return ''
		},
	},
}
</script>

<style lang="scss" scoped>
@import '../custom.scss';
.ability {
	button {
		@extend .btn-sm;
	}
}
</style>
