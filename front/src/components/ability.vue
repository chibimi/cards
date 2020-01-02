<template>
	<div class="ability">
		<div v-if="!update" class="row mx-0">
			<span class="col-3">
				{{ ability.name }} <br />
				<span class="vo">{{ ability.title }}</span>
			</span>
			<span class="col-8">
				{{ ability.description }}<br />
				<span class="vo">{{ vo.description }}</span>
			</span>
			<div class="col-1 px-0">
				<div class="float-right">
					<button class="btn-success mb-1" @click="update = true">Update</button>
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
			<div class="form-check-inline col-5">
				<label>Type</label>
				<select v-model="ability.type">
					<option value="0">None</option>
					<option value="1">Magic Ability</option>
					<option value="2">Battle Plan</option>
					<option value="3">Attack Type</option>
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
					<button v-if="ability.id && !newAbility" class="my-1" @click="update = false">Cancel</button>
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
	props: ['abilitiesList', 'ability_id', 'ability_type'],
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
			},
			vo: {},
			template: ItemTemplate,
			items: [],
			update: false,
			newAbility: false,
		}
	},
	methods: {
		get: function(id) {
			if (id == null) {
				return
			}
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT + `/abilities/${id}?lang=UK`)
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
					console.log(res)
					this.ability = res.data
					this.ability.type = this.ability_type
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
						this.ability = { type: 0 }
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
				this.ability = { type: 0 }
			}
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
