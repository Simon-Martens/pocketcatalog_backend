import xml.etree.ElementTree as ET
from xml.dom import minidom
import re
import sys

def create_zotero_rdf():
    rdf = ET.Element('rdf:RDF')
    rdf.set('xmlns:rdf', 'http://www.w3.org/1999/02/22-rdf-syntax-ns#')
    rdf.set('xmlns:z', 'http://www.zotero.org/namespaces/export#')
    rdf.set('xmlns:dc', 'http://purl.org/dc/elements/1.1/')
    rdf.set('xmlns:vcard', 'http://nwalsh.com/rdf/vCard#')
    rdf.set('xmlns:foaf', 'http://xmlns.com/foaf/0.1/')
    rdf.set('xmlns:bib', 'http://purl.org/net/biblio#')
    rdf.set('xmlns:dcterms', 'http://purl.org/dc/terms/')
    return rdf

def create_actor_element(akteur):
    if akteur.find('ORGANISATION') is not None and akteur.find('ORGANISATION').text == '1':
        actor = ET.Element('foaf:Organization')
        name = ET.SubElement(actor, 'foaf:name')
        name.text = akteur.find('NAME').text if akteur.find('NAME') is not None else "Unknown Organization"
    else:
        actor = ET.Element('foaf:Person')
        surname = ET.SubElement(actor, 'foaf:surname')
        full_name = akteur.find('NAME').text if akteur.find('NAME') is not None else "Unknown Person"
        surname.text = full_name.split(', ')[0] if ', ' in full_name else full_name
        
        if ', ' in full_name:
            given_name = ET.SubElement(actor, 'foaf:givenName')
            given_name.text = full_name.split(', ')[1]
    
    return actor

def create_memo(book_id, content, index):
    memo = ET.Element('bib:Memo')
    memo.set('rdf:about', f'#item_{book_id}_memo_{index}')
    value = ET.SubElement(memo, 'rdf:value')
    value.text = f'<div data-schema-version="9"><p>{content}</p></div>'
    return memo


def create_book_entry(band, akteure, relationships, orte, vocabular, relationships_reihen, reihen, vocabular_reihe, zotero_rdf):
    if band.find('TITEL') is None or not band.find('TITEL').text:
        return None

    book_id = band.find('ID').text if band.find('ID') is not None else "unknown"
    book = ET.Element('bib:Book')
    book.set('rdf:about', f'#item_{book_id}')

    # Add item type
    item_type = ET.SubElement(book, 'z:itemType')
    item_type.text = 'book'

    # Add title
    title = ET.SubElement(book, 'dc:title')
    title.text = band.find('TITEL').text if band.find('TITEL') is not None else "Unknown Title"

    # Add year
    year = ET.SubElement(book, 'dc:date')
    year.text = band.find('JAHR').text if band.find('JAHR') is not None else ""

    # Add editors
    editors = ET.SubElement(book, 'bib:editors')
    editors_seq = ET.SubElement(editors, 'rdf:Seq')

    # Process relationships
    for relation in relationships:
        if relation.find('BAND') is not None and relation.find('BAND').text == book_id:
            akteur_id = relation.find('AKTEUR').text if relation.find('AKTEUR') is not None else None
            beziehung_id = relation.find('BEZIEHUNG').text if relation.find('BEZIEHUNG') is not None else None
            akteur = next((a for a in akteure if a.find('ID') is not None and a.find('ID').text == akteur_id), None)
            beziehung = next((v for v in vocabular if v.find('ID') is not None and v.find('ID').text == beziehung_id), None)
            
            if akteur is not None and beziehung is not None and beziehung.find('Beziehung') is not None:
                actor_element = create_actor_element(akteur)
                
                if beziehung.find('Beziehung').text == 'wurde herausgegeben von':
                    editor_li = ET.SubElement(editors_seq, 'rdf:li')
                    editor_li.append(actor_element)
                elif beziehung.find('Beziehung').text == 'wurde verlegt von':
                    publisher = ET.SubElement(book, 'dc:publisher')
                    publisher.append(actor_element)

    # Add place of publication
    ort_id = band.find('ORTE/Value').text if band.find('ORTE/Value') is not None else None
    if ort_id:
        ort = next((o for o in orte if o.find('ID') is not None and o.find('ID').text == ort_id), None)
        if ort is not None and ort.find('NAME') is not None:
            publisher = book.find('dc:publisher')
            if publisher is None:
                publisher = ET.SubElement(book, 'dc:publisher')
                org = ET.SubElement(publisher, 'foaf:Organization')
            else:
                org = publisher.find('foaf:Organization')
                if org is None:
                    org = ET.SubElement(publisher, 'foaf:Organization')
            
            adr = ET.SubElement(org, 'vcard:adr')
            address = ET.SubElement(adr, 'vcard:Address')
            locality = ET.SubElement(address, 'vcard:locality')
            locality.text = ort.find('NAME').text

    # Add series information
    for relation in relationships_reihen:
        if relation.find('BAND') is not None and relation.find('BAND').text == book_id:
            reihe_id = relation.find('REIHE').text if relation.find('REIHE') is not None else None
            beziehung_id = relation.find('BEZIEHUNG').text if relation.find('BEZIEHUNG') is not None else None
            reihe = next((r for r in reihen if r.find('ID') is not None and r.find('ID').text == reihe_id), None)
            beziehung = next((v for v in vocabular_reihe if v.find('ID') is not None and v.find('ID').text == beziehung_id), None)
            
            if reihe is not None and beziehung is not None and reihe.find('NAME') is not None:
                series = ET.SubElement(book, 'dcterms:isPartOf')
                series_element = ET.SubElement(series, 'bib:Series')
                series_title = ET.SubElement(series_element, 'dc:title')
                series_title.text = reihe.find('NAME').text

    # Add STRUKTUR to z:numPages
    if band.find('STRUKTUR') is not None and band.find('STRUKTUR').text:
        num_pages = ET.SubElement(book, 'z:numPages')
        num_pages.text = band.find('STRUKTUR').text

    # Add description
    description = ET.SubElement(book, 'dcterms:abstract')
    desc_parts = []
    
    # Handle STATUS
    status = band.find('STATUS/Value')
    if status is not None and status.text:
        desc_parts.append(f"Status: {status.text}")
    
    # Handle NACHWEIS
    nachweis = band.find('NACHWEIS')
    if nachweis is not None and nachweis.text:
        desc_parts.append(f"Nachweis: {nachweis.text}")
    
    # Handle BIBLIO-ID
    biblio_id = band.find('BIBLIO-ID')
    if biblio_id is not None and biblio_id.text:
        desc_parts.append(f"BIBLIO-ID: {biblio_id.text}")
    
    description.text = "\n".join(desc_parts)

    # Add language
    language = ET.SubElement(book, 'z:language')
    language.text = 'Deutsch'  # Assuming the language is German

    # Add Anmerkungen as Memos
    if band.find('ANMERKUNGEN') is not None and band.find('ANMERKUNGEN').text:
        anmerkungen = band.find('ANMERKUNGEN').text
        # Split Anmerkungen by \) or )\
        memo_contents = re.split(r'\\\)|\)\\\s*', anmerkungen)
        for i, content in enumerate(memo_contents):
            if content.strip():  # Only create memo if content is not empty
                memo = create_memo(book_id, content.strip(), i)
                zotero_rdf.append(memo)
                
                # Add reference to the memo in the book entry
                is_referenced_by = ET.SubElement(book, 'dcterms:isReferencedBy')
                is_referenced_by.set('rdf:resource', f'#item_{book_id}_memo_{i}')

    return book


def transform_xml(baende_file, akteure_file, relations_file, orte_file, vocabular_file, reihen_file, relations_reihen_file, vocabular_reihe_file, output_file):
    xml_string = ""
    try:
        # Parse all XML files
        baende_tree = ET.parse(baende_file)
        akteure_tree = ET.parse(akteure_file)
        relations_tree = ET.parse(relations_file)
        orte_tree = ET.parse(orte_file)
        vocabular_tree = ET.parse(vocabular_file)
        reihen_tree = ET.parse(reihen_file)
        relations_reihen_tree = ET.parse(relations_reihen_file)
        vocabular_reihe_tree = ET.parse(vocabular_reihe_file)

        baende = baende_tree.getroot().findall('Baende')
        akteure = akteure_tree.getroot().findall('Akteure')
        relationships = relations_tree.getroot().findall('_x002A_RELATION_BaendeAkteure')
        orte = orte_tree.getroot().findall('Orte')
        vocabular = vocabular_tree.getroot().findall('_x002A_VOKABULAR_WerkeAkteure')
        reihen = reihen_tree.getroot().findall('Reihen')
        relationships_reihen = relations_reihen_tree.getroot().findall('_x002A_RELATION_BaendeReihen')
        vocabular_reihe = vocabular_reihe_tree.getroot().findall('_x002A_VOKABULAR_Reihe')

        # Create the Zotero RDF
        zotero_rdf = create_zotero_rdf()

        # Process each Band
        for band in baende:
            book_entry = create_book_entry(band, akteure, relationships, orte, vocabular, relationships_reihen, reihen, vocabular_reihe, zotero_rdf)
            if book_entry is not None:
                zotero_rdf.append(book_entry)

        # Create a pretty-printed XML string
        xml_string = ET.tostring(zotero_rdf, encoding='unicode')
        
        # For debugging, print the first 1000 characters of the XML string
        print("First 1000 characters of XML string:")
        print(xml_string[:1000])
        
        pretty_xml = minidom.parseString(xml_string).toprettyxml(indent="  ")

        # Write to output file
        with open(output_file, 'w', encoding='utf-8') as f:
            f.write(pretty_xml)

        print(f"Successfully wrote output to {output_file}")

    except Exception as e:
        print(f"An error occurred: {str(e)}", file=sys.stderr)
        if xml_string:
            print("XML string at the point of failure:", file=sys.stderr)
            print(xml_string[:1000], file=sys.stderr)  # Print the first 1000 characters
        else:
            print("Error occurred before XML string was created.", file=sys.stderr)
        raise

# Usage
try:
    transform_xml('Baende.xml', 'Akteure.xml', '_RELATION_BaendeAkteure.xml', 'Orte.xml', '_VOKABULAR_WerkeAkteure.xml', 'Reihen.xml', '_RELATION_BaendeReihen.xml', '_VOKABULAR_Reihe.xml', 'zotero_output.rdf')
except Exception as e:
    print(f"Script execution failed: {str(e)}", file=sys.stderr)
    sys.exit(1)